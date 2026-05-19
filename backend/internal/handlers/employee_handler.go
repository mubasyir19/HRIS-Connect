package handlers

import (
	"backend/internal/dto/requests"
	"backend/internal/services"
	"fmt"

	"github.com/gofiber/fiber/v3"
)

type employeeHandler struct {
	service services.EmployeeService
}

func NewEmployeeHandler(service services.EmployeeService) *employeeHandler {
	return &employeeHandler{service}
}

func (h *employeeHandler) AddNewEmployee(c fiber.Ctx) error {
	var req requests.CreateEmployeeRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request",
			"details": err.Error(),
		})
	}

	fmt.Printf("this is req :\n %v", req)
	employeeResponse, err := h.service.Create(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid field request",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "add new employee successfully",
		"data":    employeeResponse,
	})
}
