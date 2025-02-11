package error

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func HandleAppError(err error, context *fiber.Ctx, msg string) error {
	if err != nil {
		var err *AppError
		if errors.As(err, &err) {
			return context.Status(err.Code).JSON(err.Message)
		}
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": msg})
	}
	return nil
}
