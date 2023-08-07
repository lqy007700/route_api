package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/zxnlx/common"
	"github.com/zxnlx/route/proto/route"
	handler2 "github.com/zxnlx/route_api/handler"
	"github.com/zxnlx/route_api/proto/route_api"
	"strconv"
)

var (
	serviceHost = "host.docker.internal"
	servicePort = "8088"

	// 注册中心配置
	consulHost       = serviceHost
	consulPort int64 = 8500
)

// 注册中心
func initRegistry() registry.Registry {
	return consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			consulHost + ":" + strconv.FormatInt(consulPort, 10),
		}
	})
}

func main() {
	c := initRegistry()

	service := micro.NewService(
		micro.Server(server.NewServer(func(options *server.Options) {
			options.Advertise = serviceHost + ":" + servicePort
		})),
		micro.Name("go.micro.api.svcApi"),
		micro.Version("latest"),
		micro.Address(":"+servicePort),

		micro.Registry(c),
	)

	service.Init()

	svcService := route.NewRouteService("go.micro.service.route", service.Client())

	err := route_api.RegisterRouteApiHandler(service.Server(), &handler2.RouteApi{
		RouteService: svcService,
	})
	if err != nil {
		return
	}

	err = service.Run()
	if err != nil {
		common.Fatal(err)
		return
	}
}
