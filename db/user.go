package db

import "log"

type User struct {
	Id           uint64 `gorm:"autoIncrement;primaryKey;"`
	ChatId       int64  `gorm:"column:chat_id;"`
	Url          string `gorm:"column:url"`
	TaskId       uint   `gorm:"column:task_id"`
	PrevPullTime int64  `gorm:"column:prev_send_time"`
}

func CreateUser(user *User) {
	Db.Create(user)
}

func GetUsersByTaskIds(taskIds []uint) []*User {
	users := make([]*User, 0)
	Db.Where("task_id IN ?", taskIds).Find(&users)
	return users
}

func UpdateUsersPrevPullTime(user []*User) {
	if len(user) == 0 {
		log.Println("UpdateUsersPrevPullTime get a empty parameter")
		return
	}
	Db.Save(user)
}

func HasThisUrlWithTheChatId(chatId int64, url string) bool {
	var num int64
	Db.Model(&User{}).Where("chat_id = ?", chatId).Where("url = ?", url).Count(&num)
	return num > 0
}
