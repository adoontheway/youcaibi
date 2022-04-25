package main

import (
	"net/http"
	"youcaibi/api/handler"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandler() *httprouter.Router {
	r := httprouter.New()
	//user related
	r.POST("/user", handler.CreateUser)
	r.POST("/user/:username", handler.LoginUser)
	r.GET("/user/:username/videos", handler.AllUserVideo)
	r.GET("/user/:username/videos/:vid", handler.UserVideo)
	r.DELETE("/user/:username/videos/:vid", handler.DeleteUserVideo)

	// video related
	r.GET("/videos/:vid/comments", handler.VideoComments)
	r.POST("/videos/:vid/comments", handler.PostVideoComment)
	r.DELETE("/videos/:vid/comments/:cid", handler.DeleteVideoComment)
	return r
}

func main() {
	r := RegisterHandler()
	mh := NewMiddlewareHandler(r)
	http.ListenAndServe(":8000", mh)
}
