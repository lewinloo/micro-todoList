package ms

import (
	"gateway/global"
	"gateway/services"
)

var (
	User services.UserService
)

func init() {
	// 微服务实例
	User = services.NewUserService("user", global.UserClient)
}
