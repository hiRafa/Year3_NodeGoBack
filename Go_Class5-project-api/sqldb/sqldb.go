package sqldb

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {

	db, err := sql.Open("sqlite", "api.sql")
	if err != nil {
		panic("Database could not connect: " + err.Error())
	}

	DB = db

	err = createTables()
	if err != nil {
		panic("Database could not connect: " + err.Error())
	}

	fmt.Println("Tables created successfully!")
}

func createTables() error {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            email TEXT NOT NULL UNIQUE,
            password TEXT NOT NULL
        )
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table: " + err.Error())
	}

	createEventsTable := `
        CREATE TABLE IF NOT EXISTS events (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            description TEXT NOT NULL,
            location TEXT NOT NULL,
            dateTime DATETIME NOT NULL,
            user_id INTEGER,
			FOREIGN KEY (user_id) REFERENCES users(id)
        )
    `

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table: " + err.Error())
	}
	return err
}
