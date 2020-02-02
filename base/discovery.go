package base

import "net"

// Discovery is a type of object that will be able to provide seed nodes to the
// underlying peer-to-peer network.
type Discovery interface {
	// GetPeers can return any number of peers (typically less than 100)
	// and can take time to run. It can be called multiple times if, for
	// example, none of the previously returned peers were usable.
	GetPeers() ([]*net.TCPAddr, error)
}

// DiscoveryAnnouncable is implemented by Discovery objects that also support
// announcing nodes.
type DiscoveryAnnouncable interface {
	// AnnounceSelf should register the current node on the underlying
	// discovery protocol so other nodes can find it.
	AnnounceSelf(*net.TCPAddr) error
}

// DiscoveryFixed type can be used when using fixed seed nodes, such as what
// Bitcoin and other cryptocurrencies do.
type DiscoveryFixed []*net.TCPAddr

func (d DiscoveryFixed) GetPeers() ([]*net.TCPAddr, error) {
	return d, nil
}
