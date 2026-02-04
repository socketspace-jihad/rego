package http_rego

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	WrongReturnCodeErr = errors.New("http response code is wrong")
)

func Test_HTTPRegoHandler(t *testing.T) {
	r := StartRouter()
	t.Run("test GET /handle endpoint", func(t *testing.T) {

		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/handle", nil)
		if err != nil {
			t.Error(err)
			return
		}

		r.ServeHTTP(rr, req)

		if status := rr.Code; status == http.StatusOK {
			t.Error(WrongReturnCodeErr)
			return
		}
	})

	t.Run("test POST /handle endpoint", func(t *testing.T) {
		//test nil body
		req, err := http.NewRequest("POST", "/handle", nil)
		if err != nil {
			t.Error(err)
			return
		}
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)
		log.Println(rr.Code)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Error(WrongReturnCodeErr)
		}
	})

	t.Run("CRUD test", func(t *testing.T) {
		rr := httptest.NewRecorder()

		//get -> 404
		req, err := http.NewRequest("GET", "/handle", nil)
		if err != nil {
			t.Error(err)
			return
		}
		r.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusNotFound {
			t.Error(WrongReturnCodeErr)
			return
		}
		//post -> 200

		rr = httptest.NewRecorder()
		kv := KeyValue{
			Key:   "test",
			Value: "value",
		}

		data, err := json.Marshal(kv)
		if err != nil {
			t.Error(err)
			return
		}
		buffer := bytes.NewBuffer(data)

		req, err = http.NewRequest("POST", "/handle", buffer)
		if err != nil {
			t.Error(err)
			return
		}
		r.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Error(WrongReturnCodeErr)
			return
		}
		//get -> 200
		buffer = bytes.NewBuffer(data)
		rr = httptest.NewRecorder()
		req, err = http.NewRequest("GET", "/handle", buffer)
		if err != nil {
			t.Error(err)
			return
		}
		r.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Error(WrongReturnCodeErr)
			return
		}

		//delete -> 202
		buffer = bytes.NewBuffer(data)
		rr = httptest.NewRecorder()
		req, err = http.NewRequest("DELETE", "/handle", buffer)
		if err != nil {
			t.Error(err)
			return
		}
		r.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusAccepted {
			t.Error(WrongReturnCodeErr)
			return
		}
		//get -> 404
		buffer = bytes.NewBuffer(data)
		rr = httptest.NewRecorder()
		req, err = http.NewRequest("GET", "/handle", buffer)
		if err != nil {
			t.Error(err)
			return
		}
		r.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusNotFound {
			t.Error(WrongReturnCodeErr)
			return
		}
	})
}
