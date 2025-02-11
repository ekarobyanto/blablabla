package users

import (
	"errors"
	"fmt"
	error2 "github.com/username/mentoring_study_case/error"
	"github.com/username/mentoring_study_case/util"
	"net/http"
)

type UserService struct {
	repo *UserRepo
}

func NewUserService(repo *UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) FindAll() ([]User, error) {
	return u.repo.FindAll()
}

func (u *UserService) FindByID(id int) (*User, error) {
	return u.repo.FindByID(id)
}

func (u *UserService) FindByEmail(email string) (*User, error) {
	return u.repo.FindByEmail(email)
}

func (u *UserService) Create(user *CreateUserRequest) error {
	_, err := u.repo.FindByEmail(user.Email)
	if err == nil {
		return &error2.AppError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("user with email %s already exists", user.Email),
		}
	}
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return &error2.AppError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("failed to hash password: %s", err),
		}
	}

	userEntity := User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
	}

	if err := u.repo.Create(&userEntity); err != nil {
		var appErr *error2.AppError
		if errors.As(err, &appErr) {
			return appErr
		}
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}
