package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// VideoComments get all comments by video
func VideoComments(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid")
}

// PostVideoComment post comment under video
func PostVideoComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid")
}

// DeleteVideoComment
func DeleteVideoComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid")
	cid := p.ByName("cid")
}
