package service

import "dzero/godb/model"

func GetUser() *model.User  {
	var user model.User
	user.Name="43444"
	return &user
}