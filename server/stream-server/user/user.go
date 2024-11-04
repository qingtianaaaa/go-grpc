package user

import (
	pb "ggrpc/proto/stream"
	"log"

	"google.golang.org/grpc"
)

type UserServiceImpl struct {
	pb.UnimplementedUserServiceServer
}

var users = make(map[string]*pb.User)

func (*UserServiceImpl) GetUserInfo(grpc.BidiStreamingServer[pb.GetUserRequest, pb.GetUserResponse]) error {
	return nil
}

func (*UserServiceImpl) AddUser(stream grpc.BidiStreamingServer[pb.AddUserRequest, pb.AddUserResponse]) error {
	req, err := stream.Recv()
	if err != nil {
		log.Fatalln("Recv error : ", err)
	}
	if req != nil {
		users[req.User.Name] = req.User
		err = stream.Send(&pb.AddUserResponse{
			Success: true,
			Message: "user added successfully",
		})
		if err != nil {
			return err
		}
		log.Println("--------------------------")
		log.Println(users[req.User.Name])
	}
	err = stream.Send(&pb.AddUserResponse{
		Success: false,
		Message: "add failed",
	})
	return err
}

func (*UserServiceImpl) UpdateUser(stream grpc.BidiStreamingServer[pb.UpdateUserRequest, pb.UpdateUserResponse]) error {
	req, err := stream.Recv()
	if err != nil {
		log.Fatalln("recv failed : ", err)
		return err
	}
	if req != nil {
		user := req.UpdatedUser
		if _, ok := users[user.Name]; ok {
			users[user.Name] = user
			err = stream.Send(&pb.UpdateUserResponse{
				Success: true,
				Message: "user updated",
			})
		} else {
			err = stream.Send(&pb.UpdateUserResponse{
				Success: false,
				Message: "user not found",
			})
		}
		if err != nil {
			return err
		}
		log.Println("--------------------------")
		log.Println(users[user.Name])
	}
	err = stream.Send(&pb.UpdateUserResponse{
		Success: false,
		Message: "received invalid",
	})
	return err
}

func (*UserServiceImpl) DeleteUser(grpc.BidiStreamingServer[pb.DeleteUserRequest, pb.DeleteUserResponse]) error {
	return nil
}
