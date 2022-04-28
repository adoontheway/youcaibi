package util

import (
	"encoding/json"
	"io"
	"net/http"
	"youcaibi/api/defs"
)

// error response
func SendErrorResponse(w http.ResponseWriter, errRsp defs.ErrorResponse) {
	w.WriteHeader(errRsp.StatusCode)
	resStr, _ := json.Marshal(&errRsp.Error)
	io.WriteString(w, string(resStr))
}

// success response
func SendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}

// error response
func SendErrorMsgResponse(w http.ResponseWriter, code int, errMsg string) {
	w.WriteHeader(code)
	io.WriteString(w, errMsg)
}
