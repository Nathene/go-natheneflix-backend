package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const maxBytes = 1 << 20 // 1MB

type JSONResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (a *app) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) {
	out, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	if len(headers) > 0 {
		for key, val := range headers[0] {
			w.Header()[key] = val
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(out)
}

func (a *app) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) {
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)

	dec.DisallowUnknownFields()

	err := dec.Decode(data)
	if err != nil {
		log.Fatal(err)
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		log.Fatal("body must only contain a single JSON value")
	}
}

func (a *app) errorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload JSONResponse
	payload.Error = true
	payload.Message = err.Error()

	a.writeJSON(w, statusCode, payload)
}
