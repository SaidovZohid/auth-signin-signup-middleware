package utils

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("Token is Invalid")
	ErrExpiredToken = errors.New("Token is Expired")
)

type Payload struct {
	Id        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt int64     `json:"expired_at"`
}

func NewPayload(email string, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		Id:        tokenId,
		Email:     email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration).Unix(),
	}

	return payload, nil
}

func (p *Payload) Valid() error {
	if float64(time.Now().Unix()) > float64(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
