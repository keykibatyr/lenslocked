package models

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Email        string
	PasswordHash string
}

type UserService struct {
	DB *sql.DB
}

func (u *UserService) Create(email, password string) (*User, error) {
	email = strings.ToLower(email)
	hashedPaswword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("could not hash the password")
	}
	
	NewUser := User{
		Email:        email,
		PasswordHash: string(hashedPaswword),
	}


	row := u.DB.QueryRow(`
	INSERT INTO users (email, password_hash)
	VALUES ($1, $2) RETURNING id`, email, hashedPaswword)

	err = row.Scan(&NewUser.ID)
	if err != nil{
		return nil, fmt.Errorf("could not scan the query")
	}

	return &NewUser, nil
}
