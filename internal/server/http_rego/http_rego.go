package http_rego

import (
	"log"
	"net/http"
	"rego/internal/server"
)

type KeyValue struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

type HTTPRego struct{}

func StartRouter() *http.ServeMux {
	mx := http.NewServeMux()
	mx.HandleFunc("/handle", HTTPRegoHandler)

	return mx
}

func (h *HTTPRego) Serve() {
	mx := StartRouter()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mx,
	}

	log.Println("http_rego is listening on port 8080...")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

func NewHTTPRego() server.Server {
	return &HTTPRego{}
}

func init() {
	server.RegisterServer("http_rego", NewHTTPRego())
}
