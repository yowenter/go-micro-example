package main

import (
	"context"
	"fmt"

	etcd "github.com/micro/go-micro/v2/registry/etcd"



	micro "github.com/micro/go-micro/v2"
	proto "wuu.com/example/hello"
)

func main() {
	// Create a new service
	service := micro.NewService(micro.Name("greeter.client"))
	reg := etcd.NewRegistry()
	service.Init(micro.Registry(reg))
	// Initialise the client and parse command line flags
	//service.Init()

	// Create new greeter client
	greeter := proto.NewGreeterService("greeter", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.Request{Name: "WUU"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Greeting)
}
