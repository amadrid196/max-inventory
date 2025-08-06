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
	ErrRoleAlreadyAdded   = errors.New("role already added this user")
	ErrRoleNotFound       = errors.New("role not found")
	ErrInvalidPermission  = errors.New("user does not have permission to add products")
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

func (s *serv) AddUserRole(ctx context.Context, userID, roleID int64) error {
	roles, err := s.repo.GetUserRoles(ctx, userID)
	if err != nil {
		return err
	}
	for _, r := range roles {
		if r.RoleID == roleID {
			return ErrRoleAlreadyAdded
		}
	}
	return s.repo.SaveUserRole(ctx, userID, roleID)
}

func (s *serv) RemoveUserRole(ctx context.Context, userID, roleID int64) error {
	roles, err := s.repo.GetUserRoles(ctx, userID)
	if err != nil {
		return err
	}

	roleFound := false
	for _, r := range roles {
		if r.RoleID == roleID {
			roleFound = true
			break
		}
	}

	if !roleFound {
		return ErrRoleNotFound
	}

	return s.repo.RemoveUserRole(ctx, userID, roleID)
}
