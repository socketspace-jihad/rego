package http_rego

import (
	"encoding/json"
	"net/http"
	"rego/internal/core"
)

//fix the status code on the handler

func HTTPRegoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var KV KeyValue
	if r.Body != nil {
		if err := json.NewDecoder(r.Body).Decode(&KV); err != nil {
			w.Write([]byte(err.Error()))
			return
		}
	}
	switch r.Method {
	case http.MethodGet:
		val, err := core.Get(KV.Key)
		if err != nil {
			w.WriteHeader(ErrResourceNotFound.Status)
			json.NewEncoder(w).Encode(ErrResourceNotFound)
			return
		}
		KV.Value = val
		json.NewEncoder(w).Encode(KV)
	case http.MethodPost:
		if r.Body == nil {
			w.WriteHeader(ErrMalformedBody.Status)
			json.NewEncoder(w).Encode(ErrMalformedBody)
			return
		}
		if err := core.Set(KV.Key, KV.Value); err != nil {
			json.NewEncoder(w).Encode(ErrKeyExists)
			return
		}
		json.NewEncoder(w).Encode(KV)
	case http.MethodPut:
		core.Upsert(KV.Key, KV.Value)
		json.NewEncoder(w).Encode(KV)
	case http.MethodDelete:
		core.Delete(KV.Key)
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(KV)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
