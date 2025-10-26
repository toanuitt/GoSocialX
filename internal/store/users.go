package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID        int64    `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  password `json:"-"`
	CreatedAt string   `json:"created_at"`
	IsActive  bool     `json:"is_active"`
	RoleID    int64    `json:"role_id"`
	Role      Role     `json:"role"`
}

type password struct {
	text *string
	hash []byte
}

type UserStore struct {
	db *sql.DB
}

func (p *UserStore) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (username, password, email) VALUES ($1, $2, $3)
    RETURNING id, created_at
	`
	err := p.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.Email,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
