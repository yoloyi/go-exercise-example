package main

import (
	"context"
	"fmt"
	micro "github.com/micro/go-micro"
	proto "go-exercise-example/go-micro/helloworld/proto"
)

func main() {
	service := micro.NewService(micro.Name("hello.client"))
	service.Init()

	h := proto.NewHelloService("hello", service.Client())

	rsp, err := h.Hello(context.TODO(), &proto.HelloRequest{
		Name: "Yu-Yong",
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp)

}
