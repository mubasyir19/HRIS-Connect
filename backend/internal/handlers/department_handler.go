package handlers

import "backend/internal/services"

type departmentHandler struct {
	service services.DepartmentService
}

func NewDepartmentHandler(service services.DepartmentService) *departmentHandler {
	return &departmentHandler{service}
}
