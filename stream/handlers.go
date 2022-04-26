package main

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

// 一个handler就是一个goroutine?
func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid")
	vl := VIDEO_DIR + vid

	video, err := os.Open(vl)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
	defer video.Close()
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("username")
	// vid := p.ByName("vid")
	io.WriteString(w, uname)
}
