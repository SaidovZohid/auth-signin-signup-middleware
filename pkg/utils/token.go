package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(firstname, lastname, email string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(firstname, lastname, email, duration)
	if err != nil {
		return "", nil, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	return string(token), payload, err
}
