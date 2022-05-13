package controllers

import (
	"context"
	"gateway/pkg/response"
	"gateway/pkg/utils"
	"gateway/services"

	ms "gateway/micro_service"

	"github.com/gin-gonic/gin"
)

// @Tags 用户模块
// @Summary 注册接口
// @accept application/json
// @Produce application/json
// @Param params body services.UserRegisterReq true "json参数"
// @Success 200 {object} response.Base "返回体"
// @Router /register [post]
func UserRegister(c *gin.Context) {
	var userRegisterReq services.UserRegisterReq
	panicIfError(c.ShouldBindJSON(&userRegisterReq))
	res, err := ms.User.Register(context.Background(), &userRegisterReq)
	panicIfError(err)
	response.OkWithData("注册成功", gin.H{"userinfo": res.Userinfo}, c)
}

// @Tags 用户模块
// @Summary 登录接口
// @accept application/json
// @Produce application/json
// @Param params body services.UserLoginReq true "json参数"
// @Success 200 {object} response.Base "返回体"
// @Router /login [post]
func UserLogin(c *gin.Context) {
	var userLoginReq services.UserLoginReq
	panicIfError(c.ShouldBindJSON(&userLoginReq))

	res, err := ms.User.Login(context.Background(), &userLoginReq)
	panicIfError(err)

	if res.Code == 200 {
		token, err := utils.GenerateToken(uint(res.Userinfo.Id))
		panicIfError(err)
		response.OkWithData("登录成功", gin.H{"user": res.Userinfo, "token": token}, c)
	} else {
		response.Fail(int(res.Code), "用户名或密码错误", c)
	}

}
