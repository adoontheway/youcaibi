package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"youcaibi/conf"

	"github.com/julienschmidt/httprouter"
)

// 一个handler就是一个goroutine?
func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid")
	// vl := VIDEO_DIR + vid + ".mp4"

	// video, err := os.Open(vl)
	// if err != nil {
	// 	log.Printf("Error when open file :%s", err)
	// 	sendErrorResponse(w, http.StatusInternalServerError, "Internal Error") //500
	// 	return
	// }
	// w.Header().Set("Content-Type", "video/mp4")
	// http.ServeContent(w, r, "", time.Now(), video)
	// defer video.Close()

	// oss
	targetUrl := conf.GetOssAddr() + vid
	http.Redirect(w, r, targetUrl, 301)
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is too big") //400
		return
	}
	// _ FileHeader 用于文件格式验证等操作
	file, _, err := r.FormFile("file") //<form name="file"> key is file
	if err != nil {
		log.Printf("Read file stream error:%s", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error:%s", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	fileName := p.ByName("vid")
	err = ioutil.WriteFile(VIDEO_DIR+fileName, data, 0666) // no 0777 权限太大
	if err != nil {
		log.Printf("Write file error:%s", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	ossfn := "videos/" + fileName
	path := "./videos/" + fileName
	buckname := "heh-videos"
	ret := UploadToOss(ossfn, path, buckname)

	if !ret {
		log.Println("Upload file failed")
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	os.Remove(path)

	sendNormalResponse(w, "Uploaded succesfully", http.StatusCreated)
}

func testpageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("./videos/upload.html")
	t.Execute(w, nil)
}
