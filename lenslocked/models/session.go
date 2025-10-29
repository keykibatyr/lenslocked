package models

import (
	"database/sql"
	"fmt"

	"github.com/keykibatyr/lenslocked/rand"
)

type Session struct {
	Id     int
	UserID int
	//We dont stroe the raw token in th db, we need it
	//to create the user :)
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	token, err := rand.SessionToken()
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	session := Session{
		UserID: userID,
		Token:  token,
	}

	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	return nil, nil
}
