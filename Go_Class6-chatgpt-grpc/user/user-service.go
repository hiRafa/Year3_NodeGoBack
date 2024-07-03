package main

import (
	"context"

	pb "grpctest.com/go/userpb"
)

// userService implements the UserService server interface
type userService struct {
	model *UserModel
}

// GetUserById retrieves a user by ID
func (s *userService) GetUserById(ctx context.Context, req *pb.UserByIdRequest) (*pb.UserResponse, error) {
	return s.model.GetUserByID(req.Id)
}

// CreateUser creates a new user
func (s *userService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	return s.model.CreateUser(req.Name, req.Email)
}
