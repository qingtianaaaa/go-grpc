package user

import (
	"context"
	pb "ggrpc/proto"
	"log"
)

type UserServiceImpl struct {
	pb.UnimplementedUserServiceServer
}

type User struct {
	Name    string   `json:"name"`
	Age     int32    `json:"age"`
	Hobbies []string `json:"hobbies"`
	Friend  []string `json:"friend"`
}

var users = map[string]User{}

func (u *UserServiceImpl) GetUserInfo(ctx context.Context, req *pb.GetUserRequest) (resp *pb.GetUserResponse, err error) {
	if user, ok := users[req.Name]; ok {
		resp = &pb.GetUserResponse{
			User: &pb.User{
				Name:    user.Name,
				Age:     user.Age,
				Hobbies: user.Hobbies,
				Friend:  user.Friend,
			},
		}
		log.Println("--------------------------------")
		log.Println("user: "+req.Name, " has been found")
		return
	} else {
		log.Println("user not found")
		return nil, nil // or return error if you want to return error message
	}
}

func (u *UserServiceImpl) AddUser(ctx context.Context, req *pb.AddUserRequest) (resp *pb.AddUserResponse, err error) {
	pbUser := req.User
	user := User{
		Name:    pbUser.Name,
		Age:     pbUser.Age,
		Hobbies: pbUser.Hobbies,
		Friend:  pbUser.Friend,
	}
	users[pbUser.Name] = user
	resp = &pb.AddUserResponse{
		Success: true,
		Message: "add user success",
	}
	log.Println("--------------current users-----------------")
	for _, v := range users {
		log.Printf("name: %s, age: %d, hobbies: %v, friend: %v\n", v.Name, v.Age, v.Hobbies, v.Friend)
	}
	return
}

func (u *UserServiceImpl) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (resp *pb.UpdateUserResponse, err error) {
	return
}

func (u *UserServiceImpl) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (resp *pb.DeleteUserResponse, err error) {
	return
}
