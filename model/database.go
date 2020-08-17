package model

import "github.com/jinzhu/gorm"

type Database struct {
	gorm.Model
	DbName        string `json:"dbName" gorm:"type:varchar(255);not null;comment:'数据库名称'"`
	DbAlias       string `json:"dbAlias" gorm:"type:varchar(128);comment:'数据库别名'"`
	DbType        int    `json:"dbType" gorm:"comment:'数据库类型'"`
	LatestVersion string `gorm:"comment:'数据库类型'"`
	DbHost        string `gorm:"comment:'数据库地址'"`
	DbPort        int    `gorm:"comment:'端口号'"`
	ServiceName   string `gorm:"comment:'服务名称'"`
	UserName      string `gorm:"comment:'用户名称'"`
	UserPassword  string `gorm:"comment:'用户密码'"`
}
