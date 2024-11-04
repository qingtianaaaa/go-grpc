package main

import (
	pb "ggrpc/proto/stream"
	"ggrpc/server/stream-server/user"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, new(user.UserServiceImpl))
	listener, err := net.Listen("tcp", ":2345")
	if err != nil {
		log.Fatalln("error listening : ", err)
		return
	}
	err = server.Serve(listener)
	if err != nil {
		log.Fatalln("error serving : ", err)
		return
	}
}
