package server

import (
	"errors"
	"fmt"
)

var (
	servers map[string]Server = make(map[string]Server)
)

type Server interface {
	Serve()
}

type ServerFactory func() Server

func RegisterServer(name string, srv Server) {
	servers[name] = srv
}

func GetServer(name string) (Server, error) {
	srv, ok := servers[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("err: %s is not implemented yet", name))
	}
	return srv, nil
}
