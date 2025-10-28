package models

import "database/sql"

type Session struct {
	Id        int
	UserID    int
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (s *SessionService) Create(userID int) (*Session, error){
	return nil, nil
}
