package handlers

import (
	"backend/internal/dto/requests"
	"backend/internal/services"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type employeeHandler struct {
	service services.EmployeeService
}

func NewEmployeeHandler(service services.EmployeeService) *employeeHandler {
	return &employeeHandler{service}
}

func (h *employeeHandler) GetListEmployees(c fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		limit = 10
	}

	filter := make(map[string]interface{})

	employees, total, err := h.service.GetAll(filter, page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed get list employees",
		})
	}

	totalPages := (total + int64(limit) - 1) / int64(limit)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": employees,
		"meta": fiber.Map{
			"total":       total,
			"page":        page,
			"limit":       limit,
			"total_pages": totalPages,
		},
	})
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
