package db

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadVideoDeletionRecord(count int64) ([]string, error) {
	findOptions := options.Find()
	findOptions.SetLimit(count)
	cursor, err := video_del_rec_coll.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	results := make([]string, 0, 0)
	for cursor.Next(context.TODO()) {
		var result *VideoDelRec
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result.Vid)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	cursor.Close(context.TODO())
	return results, nil
}

func DelVideoDeletionRecord(vid string) error {
	deleteRes, err := video_del_rec_coll.DeleteOne(context.TODO(), bson.D{{"video_id", vid}})
	if err != nil {
		return err
	}
	if deleteRes.DeletedCount > 0 {
		return nil
	}
	return errors.New("Delete video failed")
}
