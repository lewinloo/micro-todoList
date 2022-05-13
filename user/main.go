package main

import (
	conf "user/config"
	"user/handler"
	_ "user/model"
	pb "user/services"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
)

var (
	service = "user"
	version = "latest"
)

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs(conf.ConsulAddr),
	)
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulReg),
	)
	srv.Init()

	// Register handler
	pb.RegisterUserHandler(srv.Server(), new(handler.User))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
