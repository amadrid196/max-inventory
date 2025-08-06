package service

import (
	"testing"

	"github.com/amadrid196/max-inventory/internal/entity"
	models "github.com/amadrid196/max-inventory/internal/models"
	gomock "go.uber.org/mock/gomock"
)

func TestAddProduct(t *testing.T) {
	testCases := []struct {
		UserID        int64
		Name          string
		Product       models.Products
		Email         string
		ExpectedError error
	}{
		{
			Name: "AddProduct_Success",
			Product: models.Products{
				Name:        "Product 1",
				Description: "Description 1",
				Price:       10.0,
			},
			Email:         "test@t.com",
			UserID:        int64(1),
			ExpectedError: nil,
		},
		{
			Name: "AddProduct_InvalidPermission",
			Product: models.Products{
				Name:        "Product 1",
				Description: "Description 1",
				Price:       10.0,
			},
			Email:         "test@t.com",
			UserID:        int64(2),
			ExpectedError: ErrInvalidPermission,
		},
		{
			Name: "AddProduct_EmailNotFound",
			Product: models.Products{
				Name:        "Product 1",
				Description: "Description 1",
				Price:       10.0,
			},
			Email:         "notvalid@t.com",
			UserID:        int64(-1),
			ExpectedError: ErrUserNotFound,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			ctrl, mockRepo := setupTest(t)
			defer ctrl.Finish()

			if tc.ExpectedError == nil {
				mockRepo.EXPECT().
					GetUserByEmail(gomock.Any(), tc.Email).
					Return(&entity.Users{
						ID:    tc.UserID,
						EMAIL: tc.Email,
					}, nil)
				mockRepo.EXPECT().
					GetUserRoles(gomock.Any(), gomock.Eq(tc.UserID)).
					Return([]*entity.UserRole{{UserID: tc.UserID, RoleID: 1}}, nil)

				mockRepo.EXPECT().
					SaveProduct(gomock.Any(), tc.Product.Name, tc.Product.Description, tc.Product.Price, tc.UserID).
					Return(nil)
			} else {
				if tc.ExpectedError == ErrUserNotFound {
					mockRepo.EXPECT().
						GetUserByEmail(gomock.Any(), tc.Email).
						Return(nil, ErrUserNotFound)
				} else {
					mockRepo.EXPECT().
						GetUserByEmail(gomock.Any(), tc.Email).
						Return(&entity.Users{ID: tc.UserID, EMAIL: tc.Email}, nil)

					mockRepo.EXPECT().
						GetUserRoles(gomock.Any(), gomock.Eq(tc.UserID)).
						Return([]*entity.UserRole{{UserID: tc.UserID, RoleID: 3}}, nil)
				}
			}
			s := New(mockRepo)

			err := s.AddProduct(ctx, tc.Product, tc.Email)

			if err != nil && err.Error() != tc.ExpectedError.Error() {
				t.Errorf("expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
