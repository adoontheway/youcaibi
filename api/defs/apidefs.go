package defs

// Models
type UserCredential struct {
	UserName string `json:"user_name" bson:"username"`
	Password string `json:"pwd" bson:"password"`
	Id       uint64 `json:"uid" bson:"uid"`
}

type UserInfo struct {
	UserName string `json:"user_name" bson:"username"`
	Id       uint64 `json:"uid" bson:"uid"`
}

type NewVideo struct {
	AuthorId uint64 `json:"author_id" bson:"author_id"`
	Name     string `json:"name" bson:"name"`
}

type VideoInfo struct {
	Id           string `json:"vid" bson:"vid"`
	Author       uint64 `json:"author_id" bson:"author_id"`
	Name         string `json:"name" bson:"name"`
	DisplayCtime string `json:"display_ctime" bson:"display_ctime"`
}
type CommentInfo struct {
	Id      uint64 `json:"cid" bson:"cid"`
	Uid     uint64 `json:"uid" bson:"uid"`
	Content string `json:"content" bson:"content"`
	Vid     string `json:"video_id" bson:"video_id"`
}

type SimpleSession struct {
	UserName string `json:"username" bson:"username"`
	TTL      int64  `json:"ttl" bson:"ttl"` // TTL不需要更新
	Id       string `json:"sid" bson:"sid"`
}

// Response
type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}
