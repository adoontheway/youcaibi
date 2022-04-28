package db

import (
	"youcaibi/common/db"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	// user_coll          *mongo.Collection
	// video_coll         *mongo.Collection
	// comment_coll       *mongo.Collection
	// session_coll       *mongo.Collection
	video_del_rec_coll *mongo.Collection
)

func init() {
	mongo_db, _ := db.Connect("mongodb://test:123456@localhost:7017", "video_server")
	// user_coll = mongo_db.Collection("user_coll")
	// video_coll = mongo_db.Collection("video_coll")
	// comment_coll = mongo_db.Collection("comment_coll")
	// session_coll = mongo_db.Collection("session_coll")
	video_del_rec_coll = mongo_db.Collection("video_del_rec_coll")
}
