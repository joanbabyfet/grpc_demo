package main

import (
	"context"
	pb "grpc/pb"
	"log"

	"google.golang.org/grpc"
)

const (
	// ServerAddress 连接地址
	ServerAddress string = ":8000"
)

func main() {
	// 连接服务器, grpc.WithInsecure() 跳過驗證, grpc.WithBlock() 如果没成功就不讓他往下走
	conn, err := grpc.Dial(ServerAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		//遇到致命錯誤退出，用Fatal記錄日誌後，然後程序退出
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	client := pb.NewDemoClient(conn)
	// 创建发送结构体
	req := pb.HelloRequest{
		Name: "grpc",
	}
	// 调用我们的服务(SayHello方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	res, err := client.SayHello(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call SayHello err: %v", err)
	}
	// 打印返回值
	log.Println(res)
}
