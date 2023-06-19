package token

import (
	"errors"
	"time"
)

var (
	ErrExpiredToken           = errors.New("token has expired")
	ErrInvalidToken           = errors.New("token is invalid")
	ErrTokenHasNoBearerPrefix = errors.New("token has no Bearer prefix")
)

type Payload struct {
	ID        int64  `json:"user_id"`
	Username  string `json:"username"`
	ExpiredAt int64  `json:"exp"`
}

func (payload *Payload) Valid() error {
	if time.Now().After(time.Unix(payload.ExpiredAt, 0)) {
		return ErrExpiredToken
	}
	return nil
}
