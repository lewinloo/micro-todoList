package controllers

import (
	"gateway/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Tags test
// @Summary test 接口
// @accept application/json
// @Produce application/json
// @Param page query int false "当前页第几页"
// @Param pageSize query int false "当前页的大小"
// @Success 200 {object} response.Base "返回体"
// @Router /test [get]
func Test(ctx *gin.Context) {
	// 第几页
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	// size
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "5"))

	response.OkWithData("查询成功", gin.H{
		"page":     page,
		"pageSize": pageSize,
	}, ctx)
}
