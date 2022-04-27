package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// VideoComments get all comments by video
func Comments(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// vid := p.ByName("vid")
}

// PostVideoComment post comment under video
func PostComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// vid := p.ByName("vid")
}

// DeleteVideoComment
func DeleteComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// vid := p.ByName("vid")
	// cid := p.ByName("cid")
}
