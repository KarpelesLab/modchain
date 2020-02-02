package nbitcoin

import (
	"encoding/binary"
	"errors"
	"io"
)

func varIntEncode(v uint64) []byte {
	// little endian encoding
	switch {
	case v < 0xfd:
		return []byte{byte(v)}
	case v < 0xffff:
		res := make([]byte, 3)
		res[0] = 0xfd
		binary.LittleEndian.PutUint16(res[1:], uint16(v))
		return res
	case v < 0xffffffff:
		res := make([]byte, 5)
		res[0] = 0xfe
		binary.LittleEndian.PutUint32(res[1:], uint32(v))
		return res
	default:
		res := make([]byte, 9)
		res[0] = 0xff
		binary.LittleEndian.PutUint64(res[1:], v)
		return res
	}
}

func varIntWrite(w io.Writer, v uint64) error {
	_, err := w.Write(varIntEncode(v))
	return err
}

func varIntRead(r io.Reader) (uint64, error) {
	var pfx byte
	var err error

	if rb, ok := r.(io.ByteReader); ok {
		pfx, err = rb.ReadByte()
		if err != nil {
			return 0, err
		}
	} else {
		buf := make([]byte, 1)
		_, err = r.Read(buf)
		if err != nil {
			return 0, err
		}
	}

	switch pfx {
	case 0xfd: // 16bits
		buf := make([]byte, 2)
		_, err = r.Read(buf)
		return uint64(binary.LittleEndian.Uint16(buf)), err
	case 0xfe:
		buf := make([]byte, 4)
		_, err = r.Read(buf)
		return uint64(binary.LittleEndian.Uint32(buf)), err
	case 0xff:
		buf := make([]byte, 8)
		_, err = r.Read(buf)
		return uint64(binary.LittleEndian.Uint64(buf)), err
	default: // <0xfd
		return uint64(pfx), nil
	}
}

func varStringWrite(w io.Writer, v string) error {
	return varByteWrite(w, []byte(v))
}

func varStringRead(r io.Reader, maxlen uint64) (string, error) {
	v, err := varByteRead(r, maxlen)
	if err != nil || v == nil {
		return "", err
	}

	return string(v), nil
}

func varByteWrite(w io.Writer, v []byte) error {
	err := varIntWrite(w, uint64(len(v)))
	if err != nil {
		return err
	}
	_, err = w.Write(v)
	return err
}

func varByteRead(r io.Reader, maxlen uint64) ([]byte, error) {
	l, err := varIntRead(r)
	if err != nil {
		return nil, err
	}
	if l > maxlen {
		return nil, errors.New("invalid var_str: max length exceeded")
	}
	buf := make([]byte, l)
	_, err = r.Read(buf)
	return buf, err
}
