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

	// "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"github.com/google/uuid"
)

func AddUserCredential(loginName string, pwd string) (uint64, error) {
	// stmtIns, err := dbConn.Prepare("INSERT INTO t_users (login_name, pwd) VALUES(?,?)")
	// if err != nil {
	// 	return err
	// }
	// stmtIns.Exec(loginName, pwd)
	// stmtIns.Close()
	// return nil
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
	// stmtOut, err := dbConn.Prepare("SELECT pwd FROM t_user WHERE login_name=?")
	// if err != nil {
	// 	log.Printf("%s", err)
	// 	return "", err
	// }
	// var pwd string
	// stmtOut.QueryRow(loginName).Scan(&pwd)
	// stmtOut.Close()
	// return pwd, nil
	user := &defs.UserCredential{}
	err := user_coll.FindOne(context.TODO(), bson.D{{"username", loginName}}).Decode(user)
	if err != nil {
		return "", err
	}
	return user.Password, nil
}

func DeleteUserCredential(loginName string, pwd string) error {
	// stmtDel, err := dbConn.Prepare("DELETE FROM t_users WHERE login_name = ? and pwd = ? ")
	// if err != nil {
	// 	fmt.Printf("%s", err)
	// 	return err
	// }
	// stmtDel.Exec(loginName, pwd)
	// stmtDel.Close()
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
	video_coll.InsertOne(context.TODO(), video_info)
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
