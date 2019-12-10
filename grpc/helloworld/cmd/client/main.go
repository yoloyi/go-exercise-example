package main

import (
	"fmt"
	pb "go-exercise-example/grpc/helloworld/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

const (
	address     string = "localhost:50051"
	defaultName string = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure()) // 判断 grpc service 是否活着

	// grpc.WithInsecure() 禁用传递安全型，
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn) // 创建一个 client，这里的client 是 protoc 生成的
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{
		Name: name,
	}) // 调用Service 的 say hello 方法
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Get Rpc request message", r.GetMessage())

}
