package books

import (
	"github.com/gofiber/fiber/v2"
	"github.com/username/mentoring_study_case/model"
)

type BookHandler struct {
	s *BookService
}

func NewBookHandler(s *BookService) *BookHandler {
	return &BookHandler{s: s}
}

func (h *BookHandler) GetAvailableBooks(c *fiber.Ctx) error {
	books, err := h.s.GetAllAvailableBooks()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(model.AppResponse{Code: fiber.StatusOK, Message: "Books retrieved successfully", Data: books})
}

func (h *BookHandler) CreateBook(c *fiber.Ctx) error {
	var createBookDto CreateBookDto
	if err := c.BodyParser(&createBookDto); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}
	err := h.s.Create(&createBookDto)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(model.AppResponse{Code: fiber.StatusCreated, Message: "Books requested successfully"})
}
