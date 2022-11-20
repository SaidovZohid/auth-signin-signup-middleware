package repo

import "time"

type UserStorageI interface {
	Create(user *User) error 
	Get(user_email string) (*User, error)
}

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
}
