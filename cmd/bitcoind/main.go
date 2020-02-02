package main

import (
	"log"

	"github.com/KarpelesLab/modchain/base"
	"github.com/KarpelesLab/modchain/dfixed"
	"github.com/KarpelesLab/modchain/nbitcoin"
)

func main() {
	// bitcoin protocol, blockchain, etc
	disco := dfixed.New([]string{"localhost:8333"}) // only connect to localhost to avoid messing up with public nodes
	p2p, err := nbitcoin.New("tcp", ":48333", nbitcoin.BitcoinMainMagic, nbitcoin.BitcoinVersion, &base.NetOptions{Discovery: disco})
	if err != nil {
		log.Printf("failed to initialize: %s", err)
		return
	}

	// TODO
	_ = p2p
}
