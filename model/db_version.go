package model

import "github.com/jinzhu/gorm"

type DbVersion struct {
	gorm.Model
	DbId      uint   `gorm:"comment:'数据库ID'"`
	DbVersion string `gorm:"comment:'数据库版本'"`
	Logs      string `gorm:"comment:'日志'"`
	Scripts   string `gorm:"comment:'脚本'"`
	Remark    string `gorm:"comment:'备注'"`
}
