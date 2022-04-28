package handler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"youcaibi/api/db"
	"youcaibi/api/defs"
	"youcaibi/api/session"
	"youcaibi/common/util"

	"github.com/julienschmidt/httprouter"
)

// CreateUser Create/Register
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(res, ubody); err != nil {
		util.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
	}

	if _, err := db.AddUserCredential(ubody.UserName, ubody.Password); err != nil {
		util.SendErrorResponse(w, defs.ErrorDBError)
	}

	sid := session.NewSessionId(ubody.UserName)
	su := &defs.SignedUp{
		Success:   true,
		SessionId: sid,
	}

	if resp, err := json.Marshal(su); err != nil {
		util.SendErrorResponse(w, defs.ErrorInternalErrorFault)
	} else {
		util.SendNormalResponse(w, string(resp), 200)
	}
}

// LoginUser log
func LoginUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(res, ubody); err != nil {
		util.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
	}

	uname := p.ByName("username")
	if uname != ubody.UserName {
		util.SendErrorResponse(w, defs.ErrorNotAuthed)
		return
	}
	pwd, err := db.GetUserCredential(ubody.UserName)
	if err != nil || len(pwd) == 0 || pwd != ubody.Password {
		util.SendErrorResponse(w, defs.ErrorNotAuthed)
		return
	}

	sid := session.NewSessionId(ubody.UserName)
	su := &defs.SignedUp{
		Success:   true,
		SessionId: sid,
	}

	if resp, err := json.Marshal(su); err != nil {
		util.SendErrorResponse(w, defs.ErrorInternalErrorFault)
	} else {
		util.SendNormalResponse(w, string(resp), 200)
	}
}

func GetUserInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !ValidateUser(w, r) {
		return
	}
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.NewVideo{}
	if err := json.Unmarshal(res, ubody); err != nil {
		util.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	vi, err := db.AddNewVideo(ubody.AuthorId, ubody.Name)

	if err != nil {
		util.SendErrorResponse(w, defs.ErrorDBError)
		return
	}

	if resp, err := json.Marshal(vi); err != nil {
		util.SendErrorResponse(w, defs.ErrorInternalErrorFault)
	} else {
		util.SendNormalResponse(w, string(resp), 200)
	}
}

func AddNewVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !ValidateUser(w, r) {
		return
	}
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
	if !ValidateUser(w, r) {
		return
	}
	// check user is owner
	vid := p.ByName("vid")
	err := db.DeleteVideo(vid)
	if err != nil {
		log.Printf("Error in deletevideo:%s", err)
		util.SendErrorResponse(w, defs.ErrorDBError)
		return
	}
	go util.SendDeleteVideoRequest(vid)
	util.SendNormalResponse(w, "", 204)
}
