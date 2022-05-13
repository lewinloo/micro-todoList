package handler

import (
	"context"
	"errors"
	"user/model"
	pb "user/services"
	"user/utils"

	log "go-micro.dev/v4/logger"
	"gorm.io/gorm"
)

type User struct{}

func (e *User) Login(ctx context.Context, req *pb.UserLoginReq, rsp *pb.UserInfoRes) error {
	log.Infof("Received User.Call request: %v", req)
	var user model.User
	rsp.Code = 200

	if err := model.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			rsp.Code = 400
			return nil
		}
		rsp.Code = 500
		return nil
	}

	if !user.VerifyPassword(req.Password) {
		rsp.Code = 400
		return nil
	}

	rsp.Userinfo = utils.BuildUser(user)
	return nil
}

func (e *User) Register(ctx context.Context, req *pb.UserRegisterReq, rsp *pb.UserInfoRes) error {
	if req.Password != req.PasswordConfirm {
		err := errors.New("两次密码输入不一致")
		return err
	}

	var count int64 = 0
	if err := model.DB.Model(&model.User{}).Where("username = ?", req.Username).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		err := errors.New("用户已存在")
		return err
	}

	user := model.User{
		Username: req.Username,
	}
	if err := user.SetPassword(req.Password); err != nil {
		return err
	}

	if err := model.DB.Create(&user).Error; err != nil {
		return err
	}

	rsp.Userinfo = utils.BuildUser(user)

	return nil
}
