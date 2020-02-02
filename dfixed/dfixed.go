package dfixed

import (
	"net"

	"github.com/KarpelesLab/modchain/base"
)

// discoveryFixed type can be used when using fixed seed nodes, such as what
// Bitcoin and other cryptocurrencies do. Domain names can also be returned
// and will be resolved later.
type discoveryFixed []string

// New will return a new instance of base.Discovery that will resolve to the
// list passed as argument.
func New(l []string) base.Discovery {
	return discoveryFixed(l)
}

func (d discoveryFixed) GetPeers() ([]*net.TCPAddr, error) {
	// go through the list and resolve peers
	var res []*net.TCPAddr
	var err error

	// TODO parallel resolution?
	for _, addr := range d {
		tcpa, e := net.ResolveTCPAddr("tcp", addr)
		if e != nil {
			// ignore error
			err = e
			continue
		}

		res = append(res, tcpa)
	}

	if res == nil {
		// could not find any node, return err if any
		return nil, err
	}

	return res, nil
}
