package db

import (
	"context"
	"fmt"
	"log"
)

func AddVideoDeletionRecord(vid string) error {

	video__del_info := &VideoDelRec{
		Vid: vid,
	}
	insertOne, err := video_del_rec_coll.InsertOne(context.TODO(), video__del_info)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Insert Video Delete record Successed:", insertOne.InsertedID)
	return nil
}
