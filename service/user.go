package service

import "dzero/godb/model"

func GetUser() *model.SysUser {
	var user model.SysUser
	user.Name = "43444"
	return &user
}
