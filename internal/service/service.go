package service

import (
	"context"

	"github.com/amadrid196/max-inventory/internal/models"
	"github.com/amadrid196/max-inventory/internal/repository"
)

// Service is the business logic layer for the application

//go:generate mockgen -source=service.go -destination=mock_service.go -package=service
type Service interface {
	RegisterUsers(ctx context.Context, email, name, password string) error
	LoginUsers(ctx context.Context, email, password string) (*models.Users, error)

	AddUserRole(ctx context.Context, userID, roleID int64) error
	RemoveUserRole(ctx context.Context, userID, roleID int64) error

	GetProducts(ctx context.Context) ([]models.Products, error)
	GetProduct(ctx context.Context, id int64) (*models.Products, error)
	AddProduct(ctx context.Context, products models.Products, userEmail string) error
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{repo: repo}
}
