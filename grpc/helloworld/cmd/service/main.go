package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "go-exercise-example/grpc/helloworld/proto"
)

const port = ":50051"

// 组合一个 server
type service struct {
	pb.UnimplementedGreeterServer
}

// grpc 重写的方法
// 如果不重写，会出现 error
func (s *service) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	str := fmt.Sprintf("Hello %s", in.GetName())
	log.Println(str)
	return &pb.HelloReply{Message: str}, nil
}

func main() {
	fmt.Println("-------------------")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()

	// 注册一个 grpc
	pb.RegisterGreeterServer(s, &service{})
	fmt.Println("Hello world GRpc start")
	// 等待阻塞
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve：%v", err)
	}

}
