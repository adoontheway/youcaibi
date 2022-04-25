package session

import (
	"sync"
	"time"
	"youcaibi/api/db"
	"youcaibi/api/defs"

	"github.com/google/uuid"
)

// in memory cache 足够，不需要使用reids，会增加系统复杂度
var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

func LoadSessionsFromDb() {
	m, err := db.RetrieveAllSession()
	if err != nil {
		return
	}
	m.Range(func(key, value any) bool {
		ss := value.(*defs.SimpleSession)
		sessionMap.Store(key, ss)
		return true
	})
}

func NewSessionId(username string) string {
	id := uuid.New()
	ct := nowInMilli()
	ttl := ct + 30*60*100 // session valid time: 30min
	session := &defs.SimpleSession{
		UserName: username,
		Id:       id.String(),
		TTL:      ttl,
	}
	sessionMap.Store(session.Id, session)
	db.InsertSession(session.Id, ttl, username)
	return session.Id
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := nowInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).UserName, false
	}
	return "", true
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	db.DeleteSession(sid)
}
