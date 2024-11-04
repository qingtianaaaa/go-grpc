package main

import (
	u "ggrpc/client/user"
	pb "ggrpc/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

var user u.User


func main() {
	conn,err :=grpc.Dial("localhost:1234",grpc.WithInsecure())
	if err!=nil{
		log.Fatalln("grpc.Dial failed : ",err)
		return 
	}
	defer conn.Close()
	client:=pb.NewUserServiceClient(conn)
	//增
	resp,err := client.AddUser(user.AddUserLocal(u.User{
		Name: "xiaopeng",
		Age: 12,
		Hobbies: []string{"ping-pong","badmiton"},
		Friend: []string{"aaa","ddd"},
	}))
	if err!=nil{
        log.Fatalln("client.AddUser failed : ",err)
        return 
    }
	log.Println("add user response: ", resp.Success,resp.Message)

	time.Sleep(time.Second * 2)

	//查
	getResp,err := client.GetUserInfo(user.GetUserInfoLocal("xiaopeng"))
	if err!=nil{
        log.Fatalln("client.GetUserInfo failed : ",err)
        return 
    }
	log.Println("get user response: ", getResp.User)


}