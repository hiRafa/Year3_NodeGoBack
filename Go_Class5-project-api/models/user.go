package models

import (
	"errors"

	"api.com/sqldb"
	"api.com/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := sqldb.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}


	u.ID = userId // Assign the last insert ID to u.ID
	return err
}

func (u *User) ValidateCredentials() error {
	// select only the password from the database, email is being received in the u.Email (struct populated from the request)
	// row will contain the password from the database by quering the u.Email from the SQL DB
	query := "SELECT password FROM users WHERE email = ?"
	row := sqldb.DB.QueryRow(query, u.Email)
	// extract the data
	var dbPassword string
	err := row.Scan(&dbPassword)

	if err != nil {
		return errors.New("Credentials are invalid. Please try again.")
	}

	// login password comes from the struct.
	// the struct is populated from the request, in the function login in routes/users.go
	// routes folder contains the handlers for the routes, which as basically the controllers
	loginPassword := u.Password
	passwordIsValid := utils.CheckHashPassword(dbPassword, loginPassword)

	if !passwordIsValid {
		return errors.New("Credentials are invalid. Please try again.")
	}

	return nil // no error, eveything worked fine
}

func GetAllUsers() ([]User, error) {
	query := "SELECT * FROM users"
	rows, error := sqldb.DB.Query(query)
	if error != nil {
		return nil, error
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		rows.Scan(
			&user.ID,
			&user.Email,
			&user.Password,
		)

		if error != nil {
			return nil, error
		}

		users = append(users, user)
	}

	return users, nil
}
