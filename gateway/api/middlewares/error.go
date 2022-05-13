package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 错误处理中间件
func Error() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != any(nil) {
				context.JSON(200, gin.H{
					"code":    404,
					"message": fmt.Sprintf("%s", r),
					"success": false,
				})
				context.Abort()
			}
		}()
		context.Next()
	}
}
