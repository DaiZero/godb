package util

import (
	"dzero/godb/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:NP123456@(192.168.1.18:30010)/godb?charset=utf8mb4&parseTime=True&loc=Local")
	//db, err := gorm.Open("sqlite3", "./tmp/gorm.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.SingularTable(true)
	db.AutoMigrate(&model.SysUser{},
		&model.DbSchema{},
		&model.DbConnection{},
		&model.DbTable{},
		&model.DbView{},
		&model.DbFunction{},
		&model.DbVersion{})

	var user model.SysUser
	db.First(&user, 1) // 查询id为1的User
	if len(user.Name) == 0 {
		db.Create(&model.SysUser{Name: "admin", Password: "admin", Telephone: 123456788})
	}
	defer db.Close()
	return db
}

func GetMySqlDB(dbc model.DbConnection) *gorm.DB {
	dbConnectionStr := dbc.UserName + ":" + dbc.UserPassword + "@(" + dbc.DbHost + ":"
	db, err := gorm.Open("mysql", dbConnectionStr)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	return db
}
