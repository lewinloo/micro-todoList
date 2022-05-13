package utils

import (
	"user/model"
	"user/services"
)

func BuildUser(user model.User) *services.UserModel {
	return &services.UserModel{
		Id:        uint32(user.ID),
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
		DeletedAt: user.DeletedAt.Time.Unix(),
		Username:  user.Username,
	}
}
