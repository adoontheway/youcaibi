package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"youcaibi/api/defs"

	"go.mongodb.org/mongo-driver/bson"
)

func AddUserCredential(loginName string, pwd string) error {
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
		insertOne, err := user_coll.InsertOne(context.TODO(), user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Insert User Successed:", insertOne.InsertedID)
		return nil
	}
	return errors.New("User already exists.")
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
	return user.UserName, nil
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
