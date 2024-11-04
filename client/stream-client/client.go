package main

import (
	"context"
	pb "ggrpc/proto/stream"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:2345", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("dial failed with error :", err)
		return
	}
	client := pb.NewUserServiceClient(conn)

	//增
	stream, err := client.AddUser(context.Background())
	if err != nil {
		log.Fatalln("add user failed with error :", err)
		return
	}
	err = stream.Send(&pb.AddUserRequest{User: &pb.User{
		Name:    "bob",
		Age:     12,
		Hobbies: []string{"ping-pong", "badminton"},
		Friend:  []string{"alice", "bob"},
	}})
	if err != nil {
		log.Fatalln("Send failed with error :", err)
	}
	resp, err := stream.Recv()
	if err != nil {
		log.Fatalln("receive failed with error :", err)
	}
	log.Println("received :", resp.Success, resp.Message)

	time.Sleep(5*time.Second)

	//改
	updateSteam, err := client.UpdateUser(context.Background())
	if err != nil {
		log.Fatalln("update failed with error :", err)
	}
	err = updateSteam.Send(&pb.UpdateUserRequest{UpdatedUser: &pb.User{
		Name:    "bob",
		Age:     21,
		Hobbies: []string{"read", "run"},
		Friend:  []string{"alice", "eric", "bob"},
	}})
	if err != nil {
		log.Fatalln("send failed with error :", err)
	}
	updateResp, err := updateSteam.Recv()
	if err != nil {
		log.Fatalln("recv failed with error :", err)
	}
	log.Println("----------------------------------------------------------------")
	log.Println("received :", updateResp.Success, updateResp.Message)

}
