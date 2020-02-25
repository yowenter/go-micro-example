package main

import (
	"context"
	"fmt"
	etcd "github.com/micro/go-micro/v2/registry/etcd"

	micro "github.com/micro/go-micro/v2"

	proto "wuu.com/example/hello"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Greeting = "Hello, " + req.Name
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	fmt.Println("service `Hello` starting.. ")
	service := micro.NewService(
		micro.Name("greeter"),
	)

	// Init will parse the command line flags.
	reg := etcd.NewRegistry()
	service.Init(micro.Registry(reg))
	// Register handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
