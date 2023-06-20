package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

type JWTValidator interface {
	VerifyToken(token string) (*Payload, error)
}

type JWTValidatorImpl struct {
	secretKey string
}

func NewJWTValidator(secretKey string) JWTValidator {
	return &JWTValidatorImpl{secretKey: secretKey}
}

func (maker *JWTValidatorImpl) VerifyToken(token string) (*Payload, error) {
	if !strings.HasPrefix(token, "Bearer ") {
		return nil, ErrTokenHasNoBearerPrefix
	}
	token = strings.Split(token, "Bearer ")[1]

	keyFunc := func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
