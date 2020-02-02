package dfixed

import (
	"net"

	"github.com/KarpelesLab/modchain/base"
)

// discoveryFixed type can be used when using fixed seed nodes, such as what
// Bitcoin and other cryptocurrencies do.
type discoveryFixed []*net.TCPAddr

// New will return a new instance of base.Discovery that will resolve to the
// list passed as argument.
func New(l []*net.TCPAddr) base.Discovery {
	return discoveryFixed(l)
}

func (d discoveryFixed) GetPeers() ([]*net.TCPAddr, error) {
	// Simply return the list
	return d, nil
}
