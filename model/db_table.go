package model

import "github.com/jinzhu/gorm"

type DbTable struct {
	gorm.Model
	DbId       uint   `json:"dbId" gorm:"comment:'数据库Id'"`
	TableName  string `json:"tableName" gorm:"comment:'数据库名'"`
	TableAlias string `json:"tableAlias" gorm:"comment:'数据库别名'"`
	TableDesc  string `json:"tableDesc" gorm:"comment:'数据库描述'"`
	Engine     string `json:"engine" gorm:"comment:'引擎'"`
	IsSystem   bool   `json:"isSystem" gorm:"comment:'是否为系统'"`
}
