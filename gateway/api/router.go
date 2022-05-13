package api

import (
	"gateway/api/controllers"
	"gateway/api/middlewares"
	_ "gateway/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.Cors(), middlewares.Error())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1")
	{
		v1.GET("/test", controllers.Test)
		v1.POST("/register", controllers.UserRegister)
		v1.POST("/login", controllers.UserLogin)
	}

	return router
}
