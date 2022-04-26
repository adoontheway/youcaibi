package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	user_coll          *mongo.Collection
	video_coll         *mongo.Collection
	comment_coll       *mongo.Collection
	session_coll       *mongo.Collection
	video_del_rec_coll *mongo.Collection
)

func init() {
	db, _ := connect("mongodb://test:123456@localhost:7017", "video_server")
	user_coll = db.Collection("user_coll")
	video_coll = db.Collection("video_coll")
	comment_coll = db.Collection("comment_coll")
	session_coll = db.Collection("session_coll")
	video_del_rec_coll = db.Collection("video_del_rec_coll")
}
