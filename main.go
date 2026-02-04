package main

import (
	"errors"
	"os"
	"rego/internal/server"
	_ "rego/internal/server/http_rego"
)

func main() {
	PROTOCOL := os.Getenv("REGO_PROTOCOL")
	if PROTOCOL == "" {
		panic(errors.New("you need to define REGO_PROTOCOL in your environment"))
	}

	srv, err := server.GetServer(PROTOCOL)
	if err != nil {
		panic(err)
	}

	srv.Serve()
}
