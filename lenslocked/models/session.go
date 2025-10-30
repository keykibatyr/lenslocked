package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"

	"github.com/keykibatyr/lenslocked/rand"
	
)
const (
	MinbytesperToken = 32
)
type Session struct {
	ID     int
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

	hashedToken := ss.hash(token) 

	session := Session{
		UserID: userID,
		Token:  token,
		TokenHash: hashedToken,
	}

	row := ss.DB.QueryRow(`UPDATE sessions 
	SET token_hash = $2 
	WHERE user_id = $1 
	RETURNING id`, session.UserID, session.TokenHash)
	err = row.Scan(&session.ID)
	if err == sql.ErrNoRows{
		row = ss.DB.QueryRow(`INSERT INTO sessions (user_id, token_hash) 
		VALUES ($1, $2) RETURNING id`, session.UserID, session.TokenHash)
		err = row.Scan(&session.ID)
	}
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	return nil, nil
}

func (ss *SessionService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:]) 
}
