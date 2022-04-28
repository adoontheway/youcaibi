package util

import (
	"log"
	"net/http"
	"youcaibi/common/conf"
)

func SendDeleteVideoRequest(id string) {
	addr := conf.GetLbAddr() + ":9001"
	url := "http://" + addr + "/video-delete-record/" + id
	_, err := http.Get(url)
	if err != nil {
		log.Printf("Sending deleting video request error:%s", err)
	}
}
