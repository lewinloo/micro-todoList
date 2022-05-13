package main

import (
	"gateway/api"
	"gateway/global"
	_ "gateway/global"
	_ "gateway/micro_service"
	logging "gateway/pkg/logging"
	"time"

	"go-micro.dev/v4/web"
)

// @title 微服务 API
// @version 0.0.1
// @description 技术栈 Gin + Go-Micro + Gorm + gRPC
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api/v1
func main() {

	// 开启一个 api 微服务网关 （api 网关）
	server := web.NewService(
		web.Name("httpService"),
		web.Address("127.0.0.1:4000"),
		web.Handler(api.NewRouter()),
		web.Registry(global.ConsulReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)

	if err := server.Init(); err != nil {
		logging.Error(err)
	}
	if err := server.Run(); err != nil {
		logging.Error(err)
	}
}
