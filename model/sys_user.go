package model

import "github.com/jinzhu/gorm"

type SysUser struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone uint
	Password  string
}
