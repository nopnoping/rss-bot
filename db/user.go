package db

type User struct {
	id           uint64 `gorm:"autoIncrement;primaryKey;"`
	TwitterId    int64  `gorm:"column:twitter_id"`
	Url          string `gorm:"column:url"`
	TaskId       uint   `gorm:"column:task_id"`
	PrevSendTime int64  `gorm:"column:prev_send_time"`
}

func CreateUser(user *User) {
	db.Create(user)
}

func GetUsersByTaskIds(taskIds []uint) []*User {
	users := make([]*User, 0)
	db.Where("task_id IN ?", taskIds).Find(&users)
	return users
}

func UpdateUsers(user []*User) {
	db.Save(user)
}
