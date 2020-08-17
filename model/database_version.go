package model

import "github.com/jinzhu/gorm"

type DatabaseVersion struct {
	gorm.Model
	DbId          string `gorm:"comment:'数据库ID'"`
	DbVersion     string `gorm:"comment:'数据库版本'"`
	VersionLog    string `gorm:"comment:'版本日志'"`
	VersionScript string `gorm:"comment:'版本日志'"`
	Remark        string `gorm:"comment:'备注'"`
}
