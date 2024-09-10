package main

import (
	"encoding/json"
	"net/http"
)

/*
	Encoding and decoding of JSON
    Marshalling: the act of converting a Go data structure into valid JSON.
    Unmarshalling: the act of parsing a valid JSON string into a data structure in Go.
*/

// You can use tags on struct field declarations to customize the encoded JSON key names.
type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

/*
	http.ResponseWriter is a interface and is created below
	http.Request is a struct with defined key/values like url, method, path, body etc.
*/

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker service",
	}

	out, _ := json.MarshalIndent(payload, "", "\t")

	println("json.MarshalIndent output", out)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)

}
