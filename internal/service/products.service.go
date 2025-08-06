package service

import (
	context "context"
	"errors"
	"log"

	models "github.com/amadrid196/max-inventory/internal/models"
)

var validRolesToAddProduct = []int64{1, 2}
var ErrorValidPermission = errors.New("user does not have permission to add products")

func (s *serv) GetProducts(ctx context.Context) ([]models.Products, error) {
	pp, err := s.repo.GetProducts(ctx)

	if err != nil {
		return nil, err
	}

	products := []models.Products{}
	for _, p := range pp {
		products = append(products, models.Products{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		})
	}
	return products, nil
}

func (s *serv) GetProduct(ctx context.Context, id int64) (*models.Products, error) {
	p, err := s.repo.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	products := &models.Products{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}
	return products, nil

}

func (s *serv) AddProduct(ctx context.Context, products models.Products, email string) error {

	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		log.Println("error getting user by email:", err)
		return ErrUserNotFound
	}

	roles, err := s.repo.GetUserRoles(ctx, u.ID)
	if err != nil {
		return err
	}

	userCanAdd := false

	for _, r := range roles {
		for _, vr := range validRolesToAddProduct {
			if vr == r.RoleID {
				userCanAdd = true

			}
		}
	}

	if !userCanAdd {
		return ErrorValidPermission
	}
	return s.repo.SaveProduct(
		ctx,
		products.Name,
		products.Description,
		products.Price,
		u.ID,
	)
}
