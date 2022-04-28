package handler

import (
	"net/http"
	"youcaibi/api/defs"
	"youcaibi/api/session"
	"youcaibi/common/util"
)

const (
	HEADER_FIELD_SESSION = "X-Session-Id"
	HEADER_FIELD_UNAME   = "X-User-Name"
)

func ValidateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}

	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	r.Header.Add(HEADER_FIELD_UNAME, uname)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0 {
		util.SendErrorResponse(w, defs.ErrorNotAuthed)
		return false
	}
	return true
}

// IAM 模块，用来区分用户角色和权限
// - SSO
// - RBAC role based access control
