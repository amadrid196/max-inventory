package service

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/amadrid196/max-inventory/encryption"
	"github.com/amadrid196/max-inventory/internal/entity"
	"github.com/amadrid196/max-inventory/internal/repository"
	gomock "go.uber.org/mock/gomock"
)

var ctx = context.Background()

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func setupTest(t *testing.T) (*gomock.Controller, *repository.MockRepository) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockRepository(ctrl)
	return ctrl, mockRepo
}

func mockUser(email, plainPassword string) *entity.Users {
	validPassword, _ := encryption.Encrypt([]byte(plainPassword))
	base64Password := encryption.ToBase64(validPassword)

	return &entity.Users{
		EMAIL:    email,
		PASSWORD: base64Password,
	}
}

func expectGetUserByEmail(mockRepo *repository.MockRepository, email string, found bool) {
	if found {
		storedPassword := "1234567"
		user := mockUser(email, storedPassword)
		mockRepo.EXPECT().GetUserByEmail(gomock.Any(), email).
			Return(user, nil)
	} else {
		mockRepo.EXPECT().GetUserByEmail(gomock.Any(), email).
			Return(nil, nil)
	}
}

func TestRegisterUsers(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		UserName      string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "RegisterUser_Success",
			Email:         "a@a.com",
			UserName:      "User",
			Password:      "123456",
			ExpectedError: nil,
		},
		{
			Name:          "RegisterUser_UserAlreadyExists",
			Email:         "test@t.com",
			UserName:      "User",
			Password:      "1234567",
			ExpectedError: ErrUserAlreadyExists,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			ctrl, mockRepo := setupTest(t)
			defer ctrl.Finish()

			if tc.ExpectedError == nil {
				expectGetUserByEmail(mockRepo, tc.Email, false)
				mockRepo.EXPECT().
					SaveUsers(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)
			} else {
				expectGetUserByEmail(mockRepo, tc.Email, true)
			}

			s := New(mockRepo)

			err := s.RegisterUsers(ctx, tc.Email, tc.UserName, tc.Password)

			if err != tc.ExpectedError {
				t.Errorf("expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}

}

func TestLoginUsers(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "LoginUser_Success",
			Email:         "a@a.com",
			Password:      "1234567",
			ExpectedError: nil,
		},
		{
			Name:          "LoginUser_InvalidCredentials",
			Email:         "a@a.com",
			Password:      "123456",
			ExpectedError: ErrInvalidCredentials,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			ctrl, mockRepo := setupTest(t)
			defer ctrl.Finish()

			if tc.ExpectedError == nil {
				expectGetUserByEmail(mockRepo, tc.Email, true)

			} else {
				expectGetUserByEmail(mockRepo, tc.Email, true)
			}
			s := New(mockRepo)

			_, err := s.LoginUsers(ctx, tc.Email, tc.Password)

			if err != tc.ExpectedError {
				t.Errorf("expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestUserAddUserRole(t *testing.T) {
	testCases := []struct {
		Name          string
		UserID        int64
		RoleID        int64
		ExpectedError error
	}{
		{
			Name:          "AddUserRole_Success",
			UserID:        1,
			RoleID:        2,
			ExpectedError: nil,
		},
		{
			Name:          "AddUserRole_RoleAlreadyAdded",
			UserID:        1,
			RoleID:        1,
			ExpectedError: ErrRoleAlreadyAdded,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			ctrl, mockRepo := setupTest(t)
			defer ctrl.Finish()

			mockRepo.EXPECT().
				GetUserRoles(gomock.Any(), tc.UserID).Return([]*entity.UserRole{{UserID: int64(1), RoleID: int64(1)}}, nil)

			if tc.ExpectedError == nil {
				mockRepo.EXPECT().
					SaveUserRole(gomock.Any(), tc.UserID, tc.RoleID).
					Return(nil)
			}

			s := New(mockRepo)

			err := s.AddUserRole(ctx, tc.UserID, tc.RoleID)

			if !errors.Is(err, tc.ExpectedError) {
				t.Errorf("expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestRemoveUserRole(t *testing.T) {
	testCases := []struct {
		Name          string
		UserID        int64
		RoleID        int64
		ExpectedError error
	}{
		{
			Name:          "RemoveUserRole_Success",
			UserID:        1,
			RoleID:        1,
			ExpectedError: nil,
		},
		{
			Name:          "RemoveUserRole_RoleNotFound",
			UserID:        1,
			RoleID:        3,
			ExpectedError: ErrRoleNotFound,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			ctrl, mockRepo := setupTest(t)
			defer ctrl.Finish()
			mockRepo.EXPECT().
				GetUserRoles(gomock.Any(), tc.UserID).
				Return([]*entity.UserRole{{RoleID: int64(1)}}, nil)

			if tc.ExpectedError == nil {
				mockRepo.EXPECT().
					RemoveUserRole(gomock.Any(), tc.UserID, tc.RoleID).
					Return(nil)
			}
			s := New(mockRepo)

			err := s.RemoveUserRole(ctx, tc.UserID, tc.RoleID)

			if !errors.Is(err, tc.ExpectedError) {
				t.Errorf("expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
