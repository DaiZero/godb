package util

import (
	"dzero/godb/model"
	"github.com/jinzhu/gorm"
)

func  InitDB() *gorm.DB  {
	db, err := gorm.Open("mysql", "root:123456@(192.168.13.108:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.User{})

	var user model.User
	db.First(&user, 1) // 查询id为1的User
	if len(user.Name) == 0 {
		db.Create(&model.User{Name: "dzero",Password: "123456",Telephone: 123456788})
	}
	defer db.Close()
	return db
}
