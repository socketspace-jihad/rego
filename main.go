package main

import (
	"errors"
	"os"

	_ "github.com/socketspace-jihad/rego/internal/server/http_rego"

	"github.com/socketspace-jihad/rego/internal/server"
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
