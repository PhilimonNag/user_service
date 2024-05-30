package main

import (
	"context"
	"testing"

	pb "github.com/philimonnag/user_service/proto"
)

func TestGetUserById(t *testing.T) {
	s := &userServiceServer{}

	req := &pb.GetUserByIdRequest{Id: 1}
	res, err := s.GetUserById(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if res.User.Id != 1 {
		t.Errorf("expected user ID 1, got %v", res.User.Id)
	}
}

func TestGetUsersByIds(t *testing.T) {
	s := &userServiceServer{}

	req := &pb.GetUsersByIdsRequest{Ids: []int32{1, 2}}
	res, err := s.GetUsersByIds(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(res.Users) != 2 {
		t.Errorf("expected 2 users, got %v", len(res.Users))
	}
}

func TestSearchUsers(t *testing.T) {
	s := &userServiceServer{}

	req := &pb.SearchUsersRequest{City: "LA"}
	res, err := s.SearchUsers(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(res.Users) == 0 {
		t.Errorf("expected at least 1 user, got %v", len(res.Users))
	}
}

func TestCreateUser(t *testing.T) {
	s := &userServiceServer{}

	req := &pb.CreateUserRequest{
		Fname:   "Alice",
		City:    "SF",
		Phone:   1122334455,
		Height:  5.7,
		Married: false,
	}
	res, err := s.CreateUser(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if res.User.Fname != "Alice" {
		t.Errorf("expected user name Alice, got %v", res.User.Fname)
	}
}
