package db

import (
	"context"
	"fmt"
	"testing"
	"youcaibi/api/defs"
)

// test fixture
var (
	user = defs.UserCredential{
		UserName: "123",
		Password: "456",
	}
	video = &defs.VideoInfo{
		Name: "This is a Video",
	}
)

// init -> run tests -> clear data
func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func clearTables() {
	// dbConn.Exec("truncate t_users")
	// dbConn.Exec("truncate t_video_info")
	// dbConn.Exec("truncate t_comments")
	// dbConn.Exec("truncate t_sessions")
	ctx := context.TODO()
	user_coll.Drop(ctx)
	video_coll.Drop(ctx)
	comment_coll.Drop(ctx)
	session_coll.Drop(ctx)
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("add", testAddUser)
	t.Run("get", testGetUser)
	t.Run("delete", testDeleteUser)
	t.Run("reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	uid, err := AddUserCredential(user.UserName, user.Password)
	user.Id = uid
	if err != nil {
		t.Errorf("Error on AddUser:%v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential(user.UserName)
	if err != nil {
		t.Errorf("Error on GetUser:%v", err)
	}

	if pwd != user.Password {
		t.Errorf("Error on GetUser pwd:%v", pwd)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUserCredential(user.UserName, user.Password)
	if err != nil {
		t.Errorf("Error on DeleteUser:%v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential(user.UserName)
	if err == nil {

		if pwd != "" {
			t.Errorf("Error on RegetUser: pwd is %s", pwd)
		}
	}
}

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideo)
	t.Run("GetVideo", testGetVideo)
	t.Run("GetUserVideo", testGetUserVideo)
	t.Run("DeleteVideo", testDeleteVideo)
	t.Run("RegetVideo", testRegetVideo)
}

func testAddVideo(t *testing.T) {
	t_video, err := AddNewVideo(user.Id, video.Name)
	if err != nil {
		t.Errorf("Add New video error: %s", err)
	}
	video = t_video
}
func testGetVideo(t *testing.T) {
	_, err := GetVideo(video.Id)
	if err != nil {
		t.Errorf("Get video error: %s", err)
	}
}

func testGetUserVideo(t *testing.T) {
	list, err := GetUserVideo(user.Id)
	if err != nil {
		t.Errorf("Get User video error: %s", err)
	}
	if len(list) == 0 {
		t.Error("Get Video failed")
	}
}

func testDeleteVideo(t *testing.T) {
	err := DeleteVideo(video.Id)
	if err != nil {
		t.Errorf("Delete video error: %s", err)
	}
}

func testRegetVideo(t *testing.T) {
	list, err := GetUserVideo(user.Id)
	if err != nil {
		t.Errorf("Get video by user error:%s", err)
	}
	if list != nil && len(list) != 0 {
		t.Error("Reget video failed, supposed to be empty")
	}
}

func TestCommentFlow(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddVideo", testAddVideo)
	t.Run("AddComment", testAddComment)
	t.Run("GetVideoComment", testGetVideoComment)
	t.Run("GetUserComment", testGetUserComment)
}

func testAddComment(t *testing.T) {
	err := AddNewComment(video.Id, user.Id, "This is a comment")
	if err != nil {
		t.Errorf("Add Comment error:%s", err)
	}
}

func testGetVideoComment(t *testing.T) {
	list, err := GetVideoComments(video.Id)
	if err != nil {
		t.Errorf("Get Comments error:%s", err)
	}
	if len(list) == 0 {
		t.Error("Get comment by video failed, supposed not to be empty")
	}
}

func testGetUserComment(t *testing.T) {
	list, err := GetUserComments(user.Id)
	if err != nil {
		t.Errorf("Get Comments error:%s", err)
	}
	if len(list) == 0 {
		t.Error("Get comment by video failed, supposed not to be empty")
	}
	for _, ele := range list {
		fmt.Printf("%+v\n", ele)
	}
}
