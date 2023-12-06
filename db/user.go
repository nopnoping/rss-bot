package db

import "log"

type User struct {
	Id           uint64 `gorm:"autoIncrement;primaryKey;"`
	ChatId       int64  `gorm:"column:chat_id;"`
	Url          string `gorm:"column:url"`
	TaskId       uint   `gorm:"column:task_id"`
	PrevPullTime int64  `gorm:"column:prev_send_time"`
	Title        string `gorm:"column:title"`
}

func CreateUser(user *User) {
	database.Create(user)
}

func GetUsersByTaskIds(taskIds []uint) []*User {
	users := make([]*User, 0)
	database.Where("task_id IN ?", taskIds).Find(&users)
	return users
}

func UpdateUsersPrevPullTime(user []*User) {
	if len(user) == 0 {
		log.Println("UpdateUsersPrevPullTime get a empty parameter")
		return
	}
	database.Save(user)
}

func HasThisUrlWithTheChatId(chatId int64, url string) bool {
	var num int64
	database.Model(&User{}).Where("chat_id = ?", chatId).Where("url = ?", url).Count(&num)
	return num > 0
}

func DeleteUserByChaiIdAndUrl(chatId int64, url string) {
	database.Where("chat_id = ?", chatId).Where("url = ?", url).Delete(&User{})
}

func GetUserSubscribeUrls(chatId int64) []*User {
	users := make([]*User, 0)
	database.Where("chat_id = ?", chatId).Find(&users)

	return users
}
