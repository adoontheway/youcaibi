package db

import (
	"context"
	"testing"
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
	err := AddUserCredential("123", "456")
	if err != nil {
		t.Errorf("Error on AddUser:%v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("123")
	if err != nil {
		t.Errorf("Error on GetUser:%v", err)
	}

	if pwd != "456" {
		t.Fail()
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUserCredential("123", "456")
	if err != nil {
		t.Errorf("Error on DeleteUser:%v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("123")
	if err != nil {
		t.Errorf("Error on RegetUser:%v", err)
	}

	if pwd != "" {
		t.Errorf("Error on RegetUser: pwd is %s", pwd)
	}
}
