package model

import "github.com/jinzhu/gorm"
import "time"

type DbConnection struct {
	gorm.Model
	CcName         string    `json:"dbName" gorm:"type:varchar(255);not null;comment:'连接名称'"`
	DbType         int       `json:"dbType" gorm:"comment:'数据库类型'"`
	DbVersion      string    `json:"dbVersion" gorm:"comment:'数据库版本'"`
	LatestScanTime time.Time `json:"latestScanTime" gorm:"comment:'最新扫描时间'"`
	DbHost         string    `json:"dbHost" gorm:"comment:'数据库地址'"`
	DbPort         int       `json:"dbPort" gorm:"comment:'端口号'"`
	ServiceName    string    `json:"serviceName" gorm:"comment:'服务名称'"`
	UserName       string    `json:"userName" gorm:"comment:'用户名称'"`
	UserPassword   string    `json:"userPassword" gorm:"comment:'用户密码'"`
}
