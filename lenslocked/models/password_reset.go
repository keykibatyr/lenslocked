package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/keykibatyr/lenslocked/rand"
)

const (
	DurationResetDuration = 1 * time.Hour
)

type PasswordReset struct {
	ID     int
	UserID int
	//we will store the raw token on backend not int he db
	Token     string
	TokenHash string
	ExpiresAt time.Time
}

type PasswordResetService struct {
	DB *sql.DB

	BytesPerToken int

	DurationTime time.Duration
}

func (service *PasswordResetService) Create(email string) (*PasswordReset, error) {
	email = strings.ToLower(email)
	var userID int
	row := service.DB.QueryRow(`SELECT id FROM users WHERE email = $1`, email)
	err := row.Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("could not find the user by email")
	}

	bytesPerToken := service.BytesPerToken
	if bytesPerToken < MinbytesperToken {
		bytesPerToken = MinbytesperToken
	}

	token, err := rand.ToString(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("could not create a token")
	}

	tokenHash := service.hash(token)

	duration := service.DurationTime
	if duration == 0 {
		duration = DurationResetDuration
	}

	pwReset := &PasswordReset{
		UserID:    userID,
		Token:     token,
		TokenHash: tokenHash,
		ExpiresAt: time.Now().Add(duration),
	}

	row = service.DB.QueryRow(`INSERT INTO password_resets (user_id, token_hash, expires_at)
	VALUES ($1, $2, $3) ON CONFLICT (user_id) DO UPDATE SET token_hash = $2, expires_at = $3 RETURNING id`, pwReset.UserID, pwReset.TokenHash, pwReset.ExpiresAt)
	err = row.Scan(&pwReset.ID)
	if err != nil {
		return nil, fmt.Errorf("could not insert into the table")
	}

	return pwReset, nil
}

func (service PasswordResetService) Consume(token string) (*User, error) {
	return nil, fmt.Errorf("will implement it later :)")
}

func (service PasswordResetService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}
