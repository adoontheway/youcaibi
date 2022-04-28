package db

import (
	// "database/sql"

	// _ "github.com/go-sql-driver/mysql"
	"youcaibi/common/db"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	// dbConn       *sql.DB
	// err          error
	user_coll    *mongo.Collection
	video_coll   *mongo.Collection
	comment_coll *mongo.Collection
	session_coll *mongo.Collection
)

func init() {
	// dbConn, err = sql.Open("mysql", "root:123!@#@tcp(localhost:3306)/video_server?charset=utf8")
	// if err != nil {
	// 	panic(err.Error())
	// }

	db, _ := db.Connect("mongodb://test:123456@localhost:7017", "video_server")
	user_coll = db.Collection("user_coll")
	video_coll = db.Collection("video_coll")
	comment_coll = db.Collection("comment_coll")
	session_coll = db.Collection("session_coll")

}
