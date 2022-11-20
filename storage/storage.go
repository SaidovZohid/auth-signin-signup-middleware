package storage

import (
	"github.com/SaidovZohid/auth-signin-signup-middleware/storage/postgres"
	"github.com/SaidovZohid/auth-signin-signup-middleware/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	User() repo.UserStorageI
}

type storagePg struct {
	userRepo repo.UserStorageI
}

func NewStorageI(db *sqlx.DB) StorageI {
	return &storagePg{
		userRepo: postgres.NewUser(db),
	}
}

func (s *storagePg) User() repo.UserStorageI {
	return s.userRepo
}