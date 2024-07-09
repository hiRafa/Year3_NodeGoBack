package main

import (
	"database/sql"

	_ "modernc.org/sqlite"

	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	associationspb "my-grpc-project/associationspb"
)

var DB *sql.DB

func main() {
	// Connect to the database (example using SQLite)
	db, err := sql.Open("sqlite", "api.sql")
	if err != nil {
		panic("Database could not connect: " + err.Error())
	}

	DB = db

	// err = createTables()
	// if err != nil {
	// 	panic("Database could not connect: " + err.Error())
	// }

	// fmt.Println("Tables created successfully!")

	// Initialize the gRPC server
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server instance
	server := grpc.NewServer()

	// Initialize an instance of your gRPC service implementation
	AssociationsDBModel := &AssociationsDBModel{db: db}

	// Register the UserService server with the gRPC server
	associationspb.RegisterAssociationsDBServer(server, AssociationsDBModel)

	// Register reflection service on gRPC server
	reflection.Register(server)

	log.Println("gRPC server running on port 50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// func createTables() error {
// 	createUsersTable := `
// 		CREATE TABLE IF NOT EXISTS users (
//             id INTEGER PRIMARY KEY AUTOINCREMENT,
//             email TEXT NOT NULL UNIQUE,
//             password TEXT NOT NULL
//         )
// 	`
// 	_, err := DB.Exec(createUsersTable)
// 	if err != nil {
// 		panic("Could not create users table: " + err.Error())
// 	}

// 	createEventsTable := `
//         CREATE TABLE IF NOT EXISTS events (
//             id INTEGER PRIMARY KEY AUTOINCREMENT,
//             name TEXT NOT NULL,
//             description TEXT NOT NULL,
//             location TEXT NOT NULL,
//             dateTime DATETIME NOT NULL,
//             user_id INTEGER,
// 			FOREIGN KEY (user_id) REFERENCES users(id)
//         )
//     `

// 	_, err = DB.Exec(createEventsTable)
// 	if err != nil {
// 		panic("Could not create events table: " + err.Error())
// 	}

// 	createRegistrationsTable := `
// 	CREATE TABLE IF NOT EXISTS registrations {
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		event_id INTEGER,
// 		user_id INTEGER,
// 		FOREIGN KEY(event_id) REFERENCES events(id),
// 		FOREIGN KEY(user_id) REFERENCES users(id),
// 	}
// 	`
// 	_, err = DB.Exec(createRegistrationsTable)
// 	if err != nil {
// 		panic("Could not create registrations table: " + err.Error())
// 	}

// 	return err
// }
