package models

import (
	"database/sql"
	"fmt"

	"github.com/keykibatyr/lenslocked/rand"
)
const (
	MinbytesperToken = 32
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

	BytesPerToken int
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	bytesPerToken := ss.BytesPerToken
	if ss.BytesPerToken < MinbytesperToken {
		bytesPerToken = MinbytesperToken
	}
	token, err := rand.ToString(bytesPerToken)
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
