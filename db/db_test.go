package db

import "testing"

func TestDb(t *testing.T) {
	user := &User{
		TwitterId: "123131",
		Url:       "www.luexp.com",
		TaskId:    231,
	}
	db.Create(user)
}
