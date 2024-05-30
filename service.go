package main

import (
	"context"
	"errors"
	"sync/atomic"

	pb "github.com/philimonnag/user_service/proto"
)

type User struct {
	ID      int32
	Fname   string
	City    string
	Phone   int64
	Height  float32
	Married bool
}
type userServiceServer struct {
	pb.UnimplementedUserServiceServer
}

var (
	users = []User{
		{ID: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
		{ID: 2, Fname: "John", City: "NY", Phone: 9876543210, Height: 5.9, Married: false},
	}
	userIDCounter int32 = 3
)

func (s *userServiceServer) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	for _, user := range users {
		if user.ID == req.Id {
			return &pb.GetUserByIdResponse{User: &pb.User{
				Id:      user.ID,
				Fname:   user.Fname,
				City:    user.City,
				Phone:   user.Phone,
				Height:  user.Height,
				Married: user.Married,
			}}, nil
		}
	}
	return nil, errors.New("user not found")
}

func (s *userServiceServer) GetUsersByIds(ctx context.Context, req *pb.GetUsersByIdsRequest) (*pb.GetUsersByIdsResponse, error) {
	var responseUsers []*pb.User
	for _, id := range req.Ids {
		for _, user := range users {
			if user.ID == id {
				responseUsers = append(responseUsers, &pb.User{
					Id:      user.ID,
					Fname:   user.Fname,
					City:    user.City,
					Phone:   user.Phone,
					Height:  user.Height,
					Married: user.Married,
				})
			}
		}
	}
	return &pb.GetUsersByIdsResponse{Users: responseUsers}, nil
}

func (s *userServiceServer) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
	var responseUsers []*pb.User
	for _, user := range users {
		if (req.City == "" || user.City == req.City) &&
			(req.Phone == 0 || user.Phone == req.Phone) &&
			(!req.Married || user.Married == req.Married) {
			responseUsers = append(responseUsers, &pb.User{
				Id:      user.ID,
				Fname:   user.Fname,
				City:    user.City,
				Phone:   user.Phone,
				Height:  user.Height,
				Married: user.Married,
			})
		}
	}
	return &pb.SearchUsersResponse{Users: responseUsers}, nil
}

func (s *userServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	newUser := User{
		ID:      atomic.AddInt32(&userIDCounter, 1),
		Fname:   req.Fname,
		City:    req.City,
		Phone:   req.Phone,
		Height:  req.Height,
		Married: req.Married,
	}
	users = append(users, newUser)
	return &pb.CreateUserResponse{User: &pb.User{
		Id:      newUser.ID,
		Fname:   newUser.Fname,
		City:    newUser.City,
		Phone:   newUser.Phone,
		Height:  newUser.Height,
		Married: newUser.Married,
	}}, nil
}
