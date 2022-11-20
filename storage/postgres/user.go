package postgres

import (
	"github.com/SaidovZohid/auth-signin-signup-middleware/storage/repo"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) repo.UserStorageI {
	return &userRepo{
		db: db,
	}
}

func (ur *userRepo) Create(user *repo.User) error {
	tr, err := ur.db.Begin()
	query := `
		INSERT INTO users(
			first_name,
			last_name,
			email,
			password
		) VALUES ($1, $2, $3, $4)
	`
	_, err = tr.Exec(
		query, 
		user.FirstName, 
		user.LastName, 
		user.Email, 
		user.Password,
	)
	if err != nil {
		tr.Rollback()
		return err
	}

	err = tr.Commit()
	if err != nil {
		tr.Rollback()
		return err
	}

	return nil
}

func (ur *userRepo) Get(user_email string) (*repo.User, error) {
	query := `
		SELECT 
			id,
			first_name,
			last_name,
			email,
			password,
			created_at
		FROM users WHERE email = $1
	`
	var user repo.User
	err := ur.db.QueryRow(
		query,
		user_email,
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
