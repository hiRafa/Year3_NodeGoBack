package main

import (
	"database/sql"
	"log"

	pb "grpctest.com/go/userpb"
)

// UserModel simulates a user model interacting with a database
type UserModel struct {
	DB *sql.DB // Example database connection
}

// GetUserByID retrieves a user by ID from the database
func (m *UserModel) GetUserByID(id int32) (*pb.UserResponse, error) {
	// Example query
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := m.DB.QueryRow(query, id)

	user := &pb.UserResponse{}
	err := row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return nil, err
	}

	return user, nil
}

// CreateUser inserts a new user into the database
func (m *UserModel) CreateUser(name, email string) (*pb.UserResponse, error) {
	// Example insert query
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	result, err := m.DB.Exec(query, name, email)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	id, _ := result.LastInsertId()
	user := &pb.UserResponse{
		Id:    int32(id),
		Name:  name,
		Email: email,
	}
	return user, nil
}
