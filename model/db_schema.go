package model

import "github.com/jinzhu/gorm"

type DbSchema struct {
	gorm.Model
	CcId                uint   `json:"ccId" gorm:"comment:'连接Id'"`
	Name                string `json:"dbName" gorm:"type:varchar(255);not null;comment:'数据库名称'"`
	Alias               string `json:"dbAlias" gorm:"type:varchar(255);comment:'数据库别名'"`
	Desc                string `json:"dbDesc" gorm:"type:varchar(500);not null;comment:'数据库描述'"`
	DbType              int    `json:"dbType" gorm:"comment:'数据库类型'"`
	LatestVersion       string `json:"latestVersion" gorm:"comment:'最新版本'"`
	DefaultCharacterSet string `json:"defaultCharacterSet" gorm:"type:varchar(128);comment:'默认字符集'"`
	DefaultCollation    string `json:"defaultCollation" gorm:"type:varchar(128);comment:'默认校对规则'"`
}
