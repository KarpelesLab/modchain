package nbitcoin

import (
	"fmt"
	"net"

	"github.com/KarpelesLab/modchain/base"
)

type bitcoinNet struct {
	magic     uint32 // magic value for network
	version   int32
	userAgent string

	l   *net.TCPListener
	opt *base.NetOptions
}

// New will return a new network instance for peer-to-peer communications
// complying with Bitcoin protocol.
func New(network, addr string, magic uint32, version int32, opt *base.NetOptions) (base.Network, error) {
	a, err := net.ResolveTCPAddr(network, addr)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve p2p listen address: %w", err)
	}

	l, err := net.ListenTCP(network, a)
	if err != nil {
		return nil, fmt.Errorf("failed to listen for p2p: %w", err)
	}

	res := &bitcoinNet{
		magic:     magic,
		version:   version,
		userAgent: "",
		l:         l,
		opt:       opt,
	}

	return res, nil
}
