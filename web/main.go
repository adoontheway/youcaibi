package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	r := httprouter.New()
	r.GET("/", homeHandler)
	r.POST("/", homeHandler)
	r.GET("/userhome", userHomeHandler)
	r.POST("/userhome", userHomeHandler)

	r.POST("/api", apiHandler) // api透传

	r.POST("/upload/:vid", proxyHandler)
	r.ServeFiles("/statics/*filepath", http.Dir("./template"))
	return r
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8080", r)
}
