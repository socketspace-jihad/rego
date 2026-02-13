package main

import (
	"errors"
	"os"

	"github.com/socketspace-jihad/rego/internal/server"
	_ "github.com/socketspace-jihad/rego/internal/server/grpc_rego"
	_ "github.com/socketspace-jihad/rego/internal/server/http_rego"
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
