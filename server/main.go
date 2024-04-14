package main

import (
	"context"
	"log"
	"net"

	pb "grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type DemoService struct {
	pb.UnimplementedDemoServer
}

func (s *DemoService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResp, error) {
	log.Println(req.Name)
	return &pb.HelloResp{Message: "i'm chris"}, nil
}

const (
	Address string = ":8000"
	Network string = "tcp"
)

func main() {
	// 监听本地端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")
	// 创建gRPC服务器实例
	grpcServer := grpc.NewServer()
	// 在gRPC服务器里注册我们的服务
	pb.RegisterDemoServer(grpcServer, &DemoService{})
	//加个反射 运行 grpcui -plaintext localhost:8000 才不会报错
	reflection.Register(grpcServer)
	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
