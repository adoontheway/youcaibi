package main

import (
	"net/http"
	"youcaibi/api/handler"
	"youcaibi/api/session"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandler() *httprouter.Router {
	r := httprouter.New()
	//user related
	r.POST("/user", handler.CreateUser)
	r.POST("/user/:username", handler.LoginUser)
	r.GET("/user/:username", handler.GetUserInfo)
	r.POST("/user/:username/videos", handler.AddNewVideo)
	r.GET("/user/:username/videos", handler.AllUserVideo)
	r.GET("/user/:username/videos/:vid", handler.UserVideo)
	r.DELETE("/user/:username/videos/:vid", handler.DeleteUserVideo)

	// video related
	r.GET("/videos/:vid/comments", handler.Comments)
	r.POST("/videos/:vid/comments", handler.PostComment)
	r.DELETE("/videos/:vid/comments/:cid", handler.DeleteComment)
	return r
}

func Prepare() {
	session.LoadSessionsFromDb()
}

func main() {
	Prepare()
	r := RegisterHandler()
	mh := NewMiddlewareHandler(r)
	http.ListenAndServe(":8000", mh)
}
