package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func init() {
	var err error
	if db, err = gorm.Open(sqlite.Open("rssbot.db"), &gorm.Config{}); err != nil {
		log.Fatalf("database connect err:%v\n", err)
	}

	if err = db.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("automigrate task table err:%v\n", err)
	}

	if err = db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("automigrate user table err:%v\n", err)
	}
}
