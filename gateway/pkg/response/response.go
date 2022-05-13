package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Base struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func OkWithMessage(msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Base{
		Code:    200,
		Message: msg,
		Success: true,
	})
}

func OkWithData(msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Base{
		Code:    200,
		Message: msg,
		Success: true,
		Data:    data,
	})
}

func OkWithListData(list interface{}, total int64, current int, size int, c *gin.Context) {
	data := map[string]interface{}{
		"list":    list,
		"total":   total,
		"current": current,
		"size":    size,
	}
	c.JSON(http.StatusOK, Base{
		Code:    200,
		Message: "查询成功",
		Success: true,
		Data:    data,
	})
}

func Fail(code int, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Base{
		Code:    code,
		Message: msg,
		Success: false,
	})
}
