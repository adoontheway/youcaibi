package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"youcaibi/api/defs"
)

// error response
func sendErrorResponse(w http.ResponseWriter, errRsp defs.ErrorResponse) {
	w.WriteHeader(errRsp.StatusCode)
	resStr, _ := json.Marshal(&errRsp.Error)
	io.WriteString(w, string(resStr))
}

// success response
func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
