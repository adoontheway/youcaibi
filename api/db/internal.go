package db

import (
	"context"
	"errors"
	"log"
	"sync"
	"youcaibi/api/defs"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertSession(sid string, ttl int64, username string) error {
	// ttlstr := strconv.FormatInt(ttl, 10)
	session := &defs.SimpleSession{
		UserName: username,
		TTL:      ttl,
		Id:       sid, //util.GenSonyFlake(),
	}
	_, err := session_coll.InsertOne(context.TODO(), session)
	return err
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	session := &defs.SimpleSession{}
	err := session_coll.FindOne(context.TODO(), bson.D{{"sid", sid}}).Decode(session)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func RetrieveAllSession() (*sync.Map, error) {
	cursor, err := session_coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	m := &sync.Map{}
	for cursor.Next(context.TODO()) {
		var result defs.SimpleSession
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		m.Store(result.Id, &result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	cursor.Close(context.TODO())
	return m, nil
}

func DeleteSession(sid string) error {
	deleteRes, err := session_coll.DeleteOne(context.TODO(), bson.D{{"sid", sid}})
	if err != nil {
		return err
	}
	if deleteRes.DeletedCount > 0 {
		return nil
	}
	return errors.New("Delete session failed")
}
