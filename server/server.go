package main

import (
	pb "ggrpc/proto"
	"ggrpc/server/user"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	//注册服务
	pb.RegisterUserServiceServer(server, new(user.UserServiceImpl))
	// 监听并开始 gRPC 服务器
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("error listening :1234: %v", err)
		return
	}
	server.Serve(listener)
}
