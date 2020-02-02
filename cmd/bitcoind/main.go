package main

import (
	"log"

	"github.com/KarpelesLab/modchain/base"
	"github.com/KarpelesLab/modchain/dfixed"
	"github.com/KarpelesLab/modchain/nbitcoin"
	"github.com/KarpelesLab/modchain/rjsonrpc"
)

func main() {
	// bitcoin protocol, blockchain, etc
	rpc, err := rjsonrpc.New("tcp", ":48332")
	if err != nil {
		log.Printf("failed to initialize: %s", err)
		return
	}

	disco := dfixed.New([]string{"localhost:8333"}) // only connect to localhost to avoid messing up with public nodes
	p2p, err := nbitcoin.New("tcp", ":48333", nbitcoin.BitcoinMainMagic, nbitcoin.BitcoinVersion, &base.NetOptions{Discovery: disco, RPC: rpc})
	if err != nil {
		log.Printf("failed to initialize: %s", err)
		return
	}

	// TODO
	_ = p2p
}
