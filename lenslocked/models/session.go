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
	//We dont store the raw token in th db, we need it
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

	row := ss.DB.QueryRow(`INSERT INTO sessions (user_id, token_hash)
	VALUES ($1, $2) ON CONFLICT (user_id) DO UPDATE SET token_hash = $2 RETURNING id`, session.UserID, session.TokenHash)
	err = row.Scan(&session.ID)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	return &session, nil
}

func (ss *SessionService) Delete(token string) error {
	tokenHash := ss.hash(token)

	_, err := ss.DB.Exec(`DELETE FROM sessions
	WHERE token_hash = $1`, tokenHash)
	if err != nil {
		return fmt.Errorf("deleting from db: %w", err)
	}

	return nil

}

func (ss *SessionService) User(token string) (*User, error) {
	tokenHash := ss.hash(token)

	var user User
	row := ss.DB.QueryRow(`SELECT users.id, users.email, users.password_hash FROM sessions
	INNER JOIN users ON users.id = sessions.user_id WHERE sessions.token_hash = $1`, tokenHash)
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}

	return &user, nil
}

func (ss *SessionService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:]) 
}
