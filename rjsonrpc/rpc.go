package rjsonrpc

import (
	"net"
	"net/http"

	"github.com/KarpelesLab/modchain/base"
)

type jsonRpc struct {
	s *http.Server
}

func New(network, addr string) (base.RPC, error) {
	l, err := net.Listen(network, addr)
	if err != nil {
		return nil, err
	}

	res := &jsonRpc{}

	res.s = &http.Server{
		Addr:    addr,
		Handler: res,
	}

	// TODO handle end of serve (shutdown, etc)
	go res.s.Serve(l)

	return res, nil
}

func (j *jsonRpc) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(rw, "Only POST methods are allowed", http.StatusMethodNotAllowed)
		return
	}

	// TODO: check Content-Type: is application/json
	// TODO: parse POSTed request in json: method, params, id

	// TODO: call method

	// TODO: send response as json
	rw.Write([]byte("TODO"))
}
