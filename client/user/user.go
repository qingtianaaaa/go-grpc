package user

import (
	"context"
	pb "ggrpc/proto"

	"google.golang.org/grpc"
)

type User struct {
	Name string `json:"name"`
	Age  int32    `json:"age"`
	Hobbies []string `json:"hobbies"`
	Friend []string `json:"friend"`
}

type GetUserInfo func(ctx context.Context, in *pb.GetUserRequest, opts ...grpc.CallOption) (*pb.GetUserResponse, error)  
type AddUser func(ctx context.Context,in *pb.AddUserRequest, opts ...grpc.CallOption) (*pb.AddUserResponse)
type UpdateUser func(ctx context.Context,in *pb.UpdateUserRequest, opts...grpc.CallOption) (*pb.UpdateUserResponse, error)
type DeleteUser func(ctx context.Context, in *pb.DeleteUserRequest, opts...grpc.CallOption) (*pb.DeleteUserResponse, error)



func (User)AddUserLocal(user User)(ctx context.Context, in *pb.AddUserRequest){
	return context.Background(), &pb.AddUserRequest{
        User: &pb.User{
            Name:     user.Name,
            Age:      user.Age,
            Hobbies:  user.Hobbies,
            Friend:   user.Friend,
        },
    }
}

func (User)GetUserInfoLocal(name string)(context.Context, *pb.GetUserRequest){
	return context.Background(), &pb.GetUserRequest{
        Name: name,
    }
}