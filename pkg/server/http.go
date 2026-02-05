package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type HTTPRegoConnection struct {
	*http.Client
	*http.Transport
}

type HTTPRegoConnectionConfig func(*HTTPRegoConnection)

func WithMaxIdleConnPerHost(maxConn int) HTTPRegoConnectionConfig {
	return func(hc *HTTPRegoConnection) {
		hc.Transport.MaxIdleConnsPerHost = maxConn
	}
}

func WithMaxIdleTimeout(durationSec int) HTTPRegoConnectionConfig {
	return func(hc *HTTPRegoConnection) {
		hc.Transport.IdleConnTimeout = time.Duration(durationSec * int(time.Second))
	}
}

func NewHTTPRegoConnection(configs ...HTTPRegoConnectionConfig) *HTTPRegoConnection {
	conn := &HTTPRegoConnection{
		Transport: &http.Transport{},
	}

	log.Println(configs)
	for _, config := range configs {
		if config != nil {
			config(conn)
		}
	}

	conn.Client = &http.Client{
		Transport: conn.Transport,
	}
	return conn
}

type KeyValue struct {
	Key   string
	Value any
}

// CRUD operation
func (h *HTTPRegoConnection) Get(key string) (any, error) {
	kv := KeyValue{
		Key: key,
	}
	data, err := json.Marshal(kv)
	if err != nil {
		return nil, err
	}
	buff := bytes.NewBuffer(data)
	req, err := http.NewRequest("GET", "/handle", buff)
	if err != nil {
		return nil, err
	}
	res, err := h.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	data, err = io.ReadAll(res.Body)
	json.Unmarshal(data, &kv)
	return kv.Value, nil
}

func (h *HTTPRegoConnection) Set(key string, value any) error {
	kv := KeyValue{
		Key:   key,
		Value: value,
	}

	data, err := json.Marshal(kv)
	if err != nil {
		return err
	}
	buff := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", "/handle", buff)
	if err != nil {
		return err
	}
	res, err := h.Client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("Error status code: %v", res.StatusCode)
	}

	return nil
}

func (h *HTTPRegoConnection) Delete(key string) error {
	kv := KeyValue{
		Key: key,
	}

	data, err := json.Marshal(kv)
	if err != nil {
		return err
	}
	buff := bytes.NewBuffer(data)
	req, err := http.NewRequest("DELETE", "/handle", buff)
	if err != nil {
		return err
	}
	res, err := h.Client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusAccepted {
		return fmt.Errorf("Error status code: %v", res.StatusCode)
	}

	return nil
}
