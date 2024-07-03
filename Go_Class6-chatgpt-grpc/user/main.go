package main

import (
	"database/sql"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var DB *sql.DB

func main() {
	// Connect to the database (example using SQLite)
	db, err := sql.Open("sqlite", "api.sql")
	if err != nil {
		panic("Database could not connect: " + err.Error())
	}

	DB = db

	defer db.Close()

	// Initialize the gRPC server
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server instance
	server := grpc.NewServer()

	// This line initializes an instance of your gRPC service implementation. Let's break down userService:
	userService := &userService{model: &UserModel{DB: db}}
	// Register the UserService server with the gRPC server
	pb.RegisterUserServiceServer(server, userService)

	// Register reflection service on gRPC server
	reflection.Register(server)

	log.Println("gRPC server running on port 50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
