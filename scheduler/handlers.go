package main

import (
	"net/http"
	"youcaibi/common/util"
	"youcaibi/scheduler/db"

	"github.com/julienschmidt/httprouter"
)

func vidDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid")
	if len(vid) == 0 {
		util.SendErrorMsgResponse(w, http.StatusBadRequest, "video is is not found")
		return
	}

	err := db.AddVideoDeletionRecord(vid)
	if err != nil {
		util.SendErrorMsgResponse(w, http.StatusInternalServerError, "internal server error")
		return
	}

	util.SendNormalResponse(w, "", http.StatusOK)
}
