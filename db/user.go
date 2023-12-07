package db

type User struct {
	Id           uint64 `gorm:"autoIncrement;primaryKey;"`
	ChatId       int64  `gorm:"column:chat_id;"`
	Url          string `gorm:"column:url"`
	TaskId       uint   `gorm:"column:task_id"`
	PrevSendTime int64  `gorm:"column:prev_send_time"`
	Title        string `gorm:"column:title"`
}

func CreateUser(user *User) {
	DB().Create(user)
}

func GetUsersByTaskIds(taskIds []uint) []*User {
	users := make([]*User, 0)
	DB().Where("task_id IN ?", taskIds).Find(&users)
	return users
}

func UpdateUser(user *User) {
	DB().Save(user)
}

func HasThisUrlWithTheChatId(chatId int64, url string) bool {
	var num int64
	DB().Model(&User{}).Where("chat_id = ?", chatId).Where("url = ?", url).Count(&num)
	return num > 0
}

func DeleteUserByChaiIdAndUrl(chatId int64, url string) {
	DB().Where("chat_id = ?", chatId).Where("url = ?", url).Delete(&User{})
}

func GetUserSubscribeUrls(chatId int64) []*User {
	users := make([]*User, 0)
	DB().Where("chat_id = ?", chatId).Find(&users)

	return users
}
