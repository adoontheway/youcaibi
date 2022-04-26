package taskrunner

import (
	"errors"
	"log"
	"os"
	"sync"
	"youcaibi/scheduler/db"
)

func VideoClearDispatcher(dc dataChan) error {
	res, err := db.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("Video clear dispathcer error: %v", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("no record found")
	}

	for _, id := range res {
		dc <- id
	}
	return nil
}

func deleteVideo(vid string) error {
	err := os.Remove("./videos/" + vid)
	if err != nil && !os.IsNotExist(err) {
		log.Printf("Delete video error: %v", err)
		return err
	}
	return nil
}

func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
forloop:
	for {

		select {
		case vid := <-dc:
			// 用goroutine处理的话可能会导致重复删除，但是不影响业务逻辑
			go func(id interface{}) { //goroutine 不会暂存状态，所以不能直接写string
				err := deleteVideo(vid.(string))
				if err != nil {
					errMap.Store(id, err)
					return
				}
				if err := db.DelVideoDeletionRecord(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
			}(vid)
		default:
			break forloop
		}
	}
	var err error
	errMap.Range(func(key, value any) bool {
		err = value.(error)
		if err != nil {
			return false
		}
		return true
	})
	return err
}
