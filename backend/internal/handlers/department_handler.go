package handlers

import (
	"backend/internal/dto/requests"
	"backend/internal/dto/response"
	"backend/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type departmentHandler struct {
	service services.DepartmentService
}

func NewDepartmentHandler(service services.DepartmentService) *departmentHandler {
	return &departmentHandler{service}
}

func (h *departmentHandler) GetListAllDepartment(c fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		limit = 10
	}

	filter := make(map[string]interface{})

	departments, total, err := h.service.GetAll(filter, page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed get list departments",
		})
	}

	// Convert models to response DTOs
	var departmentResponses []*response.DepartmentResponse
	for _, dept := range departments {
		departmentResponses = append(departmentResponses, response.ToDepartmentResponse(&dept))
	}

	totalPages := (total + int64(limit) - 1) / int64(limit)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": departmentResponses,
		"meta": fiber.Map{
			"total":       total,
			"page":        page,
			"limit":       limit,
			"total_pages": totalPages,
		},
	})
}

func (h *departmentHandler) AddNewDepartment(c fiber.Ctx) error {
	var req requests.CreateNewDepartment

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
			"details": err.Error(),
		})
	}

	newDepartment, err := h.service.AddNewDepartment(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Add new department successfully",
		"data":    newDepartment,
	})
}
