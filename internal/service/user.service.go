package service

import (
	"context"
	"errors"

	"github.com/amadrid196/max-inventory/encryption"
	"github.com/amadrid196/max-inventory/internal/models"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func (s *serv) RegisterUsers(ctx context.Context, email, name, password string) error {
	u, _ := s.repo.GetUserByEmail(ctx, email)
	if u != nil {
		return ErrUserAlreadyExists
	}

	bb, err := encryption.Encrypt([]byte(password))

	if err != nil {
		return err
	}

	pass := encryption.ToBase64(bb)
	return s.repo.SaveUsers(ctx, email, name, pass)
}

func (s *serv) LoginUsers(ctx context.Context, email, password string) (*models.Users, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	bb, err := encryption.FromBase64(u.PASSWORD)
	if err != nil {
		return nil, err
	}
	decryptPassword, err := encryption.Decrypt(bb)
	if err != nil {
		return nil, err
	}

	if string(decryptPassword) != password {
		return nil, ErrInvalidCredentials
	}
	return &models.Users{
		ID: u.ID, Email: u.EMAIL, Name: u.NAME}, nil
}
