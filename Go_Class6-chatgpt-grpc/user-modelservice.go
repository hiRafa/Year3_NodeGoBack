package main

import (
	"context"
	"database/sql"
	"log"

	userpb "my-grpc-project/userpb"
)

// UserModelService implements the UserServiceServer interface
type UserModelService struct {
	userpb.UnimplementedUserServiceServer
	db *sql.DB
}

// GetUserById retrieves a user by ID
func (s *UserModelService) GetUserById(ctx context.Context, req *userpb.UserByIdRequest) (*userpb.UserResponse, error) {
	// 's' is our key to the UserModelService vault
	// 's.db' is accessing a magical database tool from the vault
	// 'QueryRow' is using that tool to search for a specific treasure (user data)
	row := s.db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", req.Id)

	// Creating a new treasure chest to hold what we find
	user := &userpb.UserResponse{} /// the road pointer & is here

	// Filling our new treasure chest with the jewels (data) we found
	err := row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		// If we couldn't find the treasure or something went wrong
		return nil, err
	}

	// Returning the directions (pointer) to our new treasure chest
	return user, nil /// the road pointer was already set previously
	// return &user, nil
}

// CreateUser creates a new user
func (s *UserModelService) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.UserResponse, error) {
	// Example insert query
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	result, err := s.db.Exec(query, req.Name, req.Email)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	id, _ := result.LastInsertId()
	user := &userpb.UserResponse{
		Id:    int32(id),
		Name:  req.Name,
		Email: req.Email,
	}
	return user, nil
}
