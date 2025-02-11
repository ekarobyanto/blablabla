package auth

import (
	"github.com/gofiber/fiber/v2"
	error2 "github.com/username/mentoring_study_case/error"
	"github.com/username/mentoring_study_case/model"
)

type AuthHandler struct {
	service *AuthService
}

func NewAuthHandler(service *AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (handler *AuthHandler) Login(c *fiber.Ctx) error {
	var req LoginRequestDto
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.AppResponse{Code: fiber.StatusBadRequest, Message: err.Error()})
	}

	resp, err := handler.service.Login(req.Email, req.Password)
	if err != nil {
		return error2.HandleAppError(err, c, "Login Failed")
	}

	return c.Status(fiber.StatusOK).JSON(model.AppResponse{Code: fiber.StatusOK, Message: "Login Success", Data: resp})
}

func (handler *AuthHandler) Register(c *fiber.Ctx) error {
	var req RegisterRequestDto
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.AppResponse{Code: fiber.StatusBadRequest, Message: err.Error()})
	}

	_, err := handler.service.Register(req.Name, req.Email, req.Password)
	if err != nil {
		return error2.HandleAppError(err, c, "Register Failed")
	}

	return c.Status(fiber.StatusOK).JSON(model.AppResponse{Code: fiber.StatusOK, Message: "Register Success"})
}
