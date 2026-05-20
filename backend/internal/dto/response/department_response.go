package response

import (
	"backend/internal/models"

	"github.com/google/uuid"
)

type DepartmentResponse struct {
	Code               string     `json:"code"`
	Name               string     `json:"name"`
	HeadOfDepartmentID *uuid.UUID `json:"headOfDepartmentId"`
	ParentDepartmentID *uuid.UUID `json:"parentDepartmentId"`
	BudgetCode         string     `json:"budgetCode"`
}

func ToDepartmentResponse(department *models.Department) *DepartmentResponse {
	if department == nil {
		return nil
	}

	response := &DepartmentResponse{
		Code:               department.Code,
		Name:               department.Name,
		HeadOfDepartmentID: department.HeadOfDepartmentID,
		ParentDepartmentID: department.ParentDepartmentID,
		BudgetCode:         department.BudgetCode,
	}

	return response
}
