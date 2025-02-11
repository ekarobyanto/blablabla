package auth

import (
	"errors"
	error2 "github.com/username/mentoring_study_case/error"
	"github.com/username/mentoring_study_case/model"
	"github.com/username/mentoring_study_case/users"
	"github.com/username/mentoring_study_case/util"
	"net/http"
	"strconv"
	"time"
)

type AuthService struct {
	userService *users.UserService
}

func NewAuthService(userService *users.UserService) *AuthService {
	return &AuthService{userService: userService}
}

func (a *AuthService) Login(email, password string) (model.AppResponse, error) {
	user, err := a.userService.FindByEmail(email)
	if err != nil {
		var appErr *error2.AppError
		if errors.As(err, &appErr) {
			return model.AppResponse{}, appErr
		}
		return model.AppResponse{}, &error2.AppError{Message: "user not exist", Code: http.StatusNotFound}
	}

	if !util.CheckPassword(user.Password, password) {
		return model.AppResponse{}, &error2.AppError{Message: "invalid password", Code: http.StatusBadRequest}
	}

	token, err := util.GenerateToken(strconv.Itoa(user.ID), time.Duration(2400)*time.Hour)
	return model.AppResponse{Code: http.StatusOK, Message: "Login Success", Data: LoginResponseDto{Token: token}}, nil
}
func (a *AuthService) Register(name, email, password string) (model.AppResponse, error) {
	_, err := a.userService.FindByEmail(email)
	if err == nil {
		return model.AppResponse{}, &error2.AppError{Message: "user already exist", Code: http.StatusBadRequest}
	}

	if err := a.userService.Create(&users.CreateUserRequest{Name: name, Email: email, Password: password}); err != nil {
		var appErr *error2.AppError
		if errors.As(err, &appErr) {
			return model.AppResponse{}, &error2.AppError{Message: appErr.Message, Code: appErr.Code}
		}
		return model.AppResponse{}, &error2.AppError{Message: "failed to create user", Code: http.StatusInternalServerError}
	}
	r, _ := a.Login(email, password)
	return model.AppResponse{Code: http.StatusOK, Message: "Register Success", Data: r.Data}, nil
}
