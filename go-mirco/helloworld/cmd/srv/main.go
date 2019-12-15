package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	proto "go-exercise-example/go-mirco/helloworld/proto"
)

type Greeter struct {
}

func (g *Greeter) Hello(ctx context.Context, in *proto.HelloRequest, out *proto.HelloReply) error {
	out.Message = "Hello " + in.Name
	return nil
}
func main() {
	srv := micro.NewService(micro.Name("hello"))
	srv.Init()
	err := proto.RegisterHelloServiceHandler(srv.Server(), new(Greeter))

	if err != nil {
		fmt.Println(err)
	}
	// 运行服务
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}


}
