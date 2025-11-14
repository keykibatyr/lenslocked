package models

import (
	"database/sql"
	"fmt"
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
	bytesPerToken := service.BytesPerToken
	if bytesPerToken < MinbytesperToken {
		bytesPerToken = MinbytesperToken
	}

	token, err := rand.ToString(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("could not create a token")
	}
	
	passRes := PasswordReset {
		Token: token,
	}

	return &passRes, fmt.Errorf("need to implement")
}

func (service PasswordResetService) Consume(token string) (*User, error) {
	return nil, fmt.Errorf("will implement it later :)")
}
