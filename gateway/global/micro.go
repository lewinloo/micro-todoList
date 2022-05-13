package global

import (
	"gateway/wrappers"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
)

var (
	UserClient client.Client
	ConsulReg  registry.Registry
)

func init() {
	// 集成consul
	consulRegistry := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	ConsulReg = consulRegistry

	// Create service
	userMicroService := micro.NewService(
		// 注册 consul
		micro.Name("user.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)
	UserClient = userMicroService.Client()
}
