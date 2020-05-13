package util

import (
	"dzero/godb/model"
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitDB() *gorm.DB {
	//db, err := gorm.Open("mysql", "root:123456@(192.168.13.108:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open("sqlite3", "./tmp/gorm.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.User{})

	var user model.User
	db.First(&user, 1) // 查询id为1的User
	if len(user.Name) == 0 {
		db.Create(&model.User{Name: "dzero", Password: "123456", Telephone: 123456788})
	}
	defer db.Close()
	return db
}
