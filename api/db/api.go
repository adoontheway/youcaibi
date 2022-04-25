package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
	"youcaibi/api/defs"
	"youcaibi/api/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/google/uuid"
)

func AddUserCredential(loginName string, pwd string) (uint64, error) {

	user := &defs.UserCredential{}
	err := user_coll.FindOne(context.TODO(), bson.D{{"username", loginName}}).Decode(user)
	if err != nil {
		user.Password = pwd
		user.UserName = loginName
		user.Id = util.GenSonyFlake()
		insertOne, err := user_coll.InsertOne(context.TODO(), user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Insert User Successed:", insertOne.InsertedID)
		return user.Id, nil
	}
	return 0, errors.New("User already exists.")
}

func GetUserCredential(loginName string) (string, error) {

	user := &defs.UserCredential{}
	err := user_coll.FindOne(context.TODO(), bson.D{{"username", loginName}}).Decode(user)
	if err != nil {
		return "", err
	}
	return user.Password, nil
}

func DeleteUserCredential(loginName string, pwd string) error {

	deleteRes, err := user_coll.DeleteOne(context.TODO(), bson.D{{"username", loginName}, {"password", pwd}})
	if err != nil {
		return err
	}
	if deleteRes.DeletedCount > 0 {
		return nil
	}
	return errors.New("Delete user failed")
}

func AddNewVideo(aid uint64, name string) (*defs.VideoInfo, error) {
	vid := uuid.New()
	t := time.Now()
	ctime := t.Format("Jan 02 2016, 15:40:05")
	video_info := &defs.VideoInfo{
		Id:           vid.String(),
		Author:       aid,
		Name:         name,
		DisplayCtime: ctime,
	}
	insertOne, err := video_coll.InsertOne(context.TODO(), video_info)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Insert User Successed:", insertOne.InsertedID)
	return video_info, nil
}

func DeleteVideo(aid uint64, vid string) error {
	deleteRes, err := video_coll.DeleteOne(context.TODO(), bson.D{{"author_id", aid}, {"vid", vid}})
	if err != nil {
		return err
	}
	if deleteRes.DeletedCount > 0 {
		return nil
	}
	return errors.New("Delete video failed")
}

// GetUserVideo TODO 分页
func GetUserVideo(aid uint64) ([]*defs.VideoInfo, error) {
	findOptions := options.Find()
	findOptions.SetLimit(10)
	cursor, err := video_coll.Find(context.TODO(), bson.D{{"author_id", aid}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	results := make([]*defs.VideoInfo, 0, 0)
	for cursor.Next(context.TODO()) {
		var result defs.VideoInfo
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, &result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	cursor.Close(context.TODO())
	return results, nil
}

func GetVideo(vid string) (*defs.VideoInfo, error) {
	video := &defs.VideoInfo{}
	err := video_coll.FindOne(context.TODO(), bson.D{{"vid", vid}}).Decode(video)
	if err != nil {
		return nil, err
	}
	return video, nil
}

// AddNewComment add new comment
func AddNewComment(vid string, uid uint64, content string) error {
	comment := defs.CommentInfo{
		Vid:     vid,
		Uid:     uid,
		Id:      util.GenSonyFlake(),
		Content: content,
	}
	insertOne, err := comment_coll.InsertOne(context.TODO(), comment)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Insert Comment Successed:", insertOne.InsertedID)
	return nil
}

// GetVideoComment get comment list by video
func GetVideoComments(vid string) ([]*defs.CommentInfo, error) {
	findOptions := options.Find()
	findOptions.SetLimit(10)
	cursor, err := comment_coll.Find(context.TODO(), bson.D{{"video_id", vid}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	results := make([]*defs.CommentInfo, 0, 0)
	for cursor.Next(context.TODO()) {
		var result defs.CommentInfo
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, &result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	cursor.Close(context.TODO())
	return results, nil
}

// GetUserComments get comments by user
func GetUserComments(uid uint64) ([]*defs.CommentInfo, error) {
	findOptions := options.Find()
	findOptions.SetLimit(10)
	cursor, err := comment_coll.Find(context.TODO(), bson.D{{"uid", uid}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	results := make([]*defs.CommentInfo, 0, 0)
	for cursor.Next(context.TODO()) {
		var result defs.CommentInfo
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, &result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	cursor.Close(context.TODO())
	return results, nil
}
