package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	
	return string(bcryptPassword), err
}

func CheckHashPassword(dbHashPassword, loginPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbHashPassword), []byte(loginPassword))
	return err == nil // returns true if there is no error. returns false if there is an error
}