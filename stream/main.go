package main

import (
	"net/http"
	"youcaibi/common/util"

	"github.com/julienschmidt/httprouter"
)

type middlewareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

func NewMiddlewareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middlewareHandler{
		r: r,
		l: NewConnLimiter(cc),
	}

	return m
}

func (m middlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		util.SendErrorMsgResponse(w, http.StatusTooManyRequests, "too many request")
		return
	}
	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func RegisterHandlers() *httprouter.Router {
	r := httprouter.New()
	r.GET("/videos/:vid", streamHandler)
	r.POST("/upload/:vid", uploadHandler)
	r.GET("/testpage", testpageHandler)
	return r
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddlewareHandler(r, 2)
	http.ListenAndServe(":9000", mh)
}
