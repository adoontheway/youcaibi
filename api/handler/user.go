package handler

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// CreateUser Create/Register
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Create User")
}

// LoginUser log
func LoginUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("username")
	io.WriteString(w, uname)
}

// AllUserVideo get all vedios by user
func AllUserVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("username")
	io.WriteString(w, uname)
}

func UserVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("username")
	// vid := p.ByName("vid")
	io.WriteString(w, uname)
}

func DeleteUserVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("username")
	// vid := p.ByName("vid")
	io.WriteString(w, uname)
}
