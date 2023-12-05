package db

type User struct {
	id        uint64 `gorm:"autoIncrement;primaryKey;"`
	TwitterId string `gorm:"column:twitter_id"`
	Url       string `gorm:"column:url"`
	TaskId    uint   `gorm:"column:task_id"`
}

func CreateUser(user *User) {
	db.Create(user)
}

func GetUsersByTaskIds(taskIds []uint) []*User {
	users := make([]*User, 0)
	db.Where("task_id IN ?", taskIds).Find(&users)
	return users
}
