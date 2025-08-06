package repository

import (
	"context"

	"github.com/amadrid196/max-inventory/internal/entity"
	"github.com/jmoiron/sqlx"
)

// Repository is an interface that defines the methods for interacting with the user repository.
//go:generate mockgen -source=repository.go -destination=mock_repository.go -package=repository

type Repository interface {
	SaveUsers(ctx context.Context, email, name, password string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.Users, error)

	SaveUserRole(ctx context.Context, userID, roleID int64) error
	RemoveUserRole(ctx context.Context, userID, roleID int64) error
	GetUserRoles(ctx context.Context, userID int64) ([]*entity.UserRole, error)
	SaveProduct(ctx context.Context, name, descripcion string, price float32, createBy int64) error
	GetProducts(ctx context.Context) ([]entity.Products, error)
	GetProduct(ctx context.Context, id int64) (*entity.Products, error)
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{db: db}
}
