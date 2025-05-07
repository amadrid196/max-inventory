package repository

import (
	"context"

	"github.com/amadrid196/max-inventory/internal/entity"
)

const (
	qryInsertUser = `INSERT INTO users (email, name, password) VALUES (?, ?, ?)`

	qryByUserEmail = `
	SELECT id, email, name, password 
	FROM users WHERE email = ?`
)

func (r *repo) SaveUsers(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser, email, name, password)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.Users, error) {
	u := &entity.Users{}
	err := r.db.GetContext(ctx, u, qryByUserEmail, email)
	return u, err
}
