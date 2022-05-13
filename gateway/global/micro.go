package global

import (
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
)

var (
	MicroClient client.Client
	ConsulReg   registry.Registry
)

func init() {
	// 集成consul
	consulRegistry := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	ConsulReg = consulRegistry

	// Create service
	srv := micro.NewService(
		// 注册 consul
		micro.Registry(consulRegistry),
	)
	MicroClient = srv.Client()
}
