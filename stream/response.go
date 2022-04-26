package main

import (
	"io"
	"net/http"
)

// error response
func sendErrorResponse(w http.ResponseWriter, code int, errMsg string) {
	w.WriteHeader(code)
	io.WriteString(w, errMsg)
}

// success response
func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
