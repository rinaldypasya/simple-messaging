package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/rinaldypasya/simple-messaging/message"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(message.Message{})
	return db
}
