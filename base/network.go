package base

type Network interface {
}

// NetOptions are options passed to network instances which contains basic
// settings needed for network.
type NetOptions struct {
	// Discovery is used for peer-to-peer networks to find other peers
	Discovery Discovery

	RPC RPC
}
