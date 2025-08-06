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

	qryInsertUserRoles = `INSERT INTO user_roles (user_id, role_id) VALUES (:user_id, :role_id)`
	qryDeleteUserRoles = `DELETE FROM user_roles WHERE user_id = :user_id AND role_id = :role_id`
)

func (r *repo) SaveUsers(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser, email, name, password)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.Users, error) {
	u := &entity.Users{}
	err := r.db.GetContext(ctx, u, qryByUserEmail, email)

	if err != nil {
		return nil, err
	}
	return u, err
}

func (r *repo) SaveUserRole(ctx context.Context, userID, roleID int64) error {
	data := entity.UserRole{
		UserID: userID,
		RoleID: roleID,
	}
	_, err := r.db.NamedExecContext(ctx, qryInsertUserRoles, data)
	return err
}

func (r *repo) RemoveUserRole(ctx context.Context, userID, roleID int64) error {
	data := entity.UserRole{
		UserID: userID,
		RoleID: roleID,
	}
	_, err := r.db.NamedExecContext(ctx, qryDeleteUserRoles, data)
	return err
}

func (r *repo) GetUserRoles(ctx context.Context, userID int64) ([]*entity.UserRole, error) {
	roles := []*entity.UserRole{}
	err := r.db.SelectContext(ctx, &roles, "select user_id, roles_id from user_roles where user_id = ?", userID)

	if err != nil {
		return nil, err
	}

	return roles, nil
}
