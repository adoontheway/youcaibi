package main

import (
	"net/http"
	"youcaibi/scheduler/taskrunner"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	r := httprouter.New()
	r.GET("/videos-delete-record/:vid", vidDelRecHandler)
	return r
}

func main() {
	go taskrunner.Start()
	r := RegisterHandlers()
	http.ListenAndServe(":9001", r) // this is  an blocking method
}
