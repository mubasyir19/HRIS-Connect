package response

import (
	"backend/internal/models"

	"github.com/google/uuid"
)

type HeadOfDepartmentInfo struct {
	ID       string `json:"id"`
	FullName string `json:"fullname"`
}

type DepartmentResponse struct {
	ID                 uuid.UUID             `json:"id"`
	Code               string                `json:"code"`
	Name               string                `json:"name"`
	HeadOfDepartmentID *uuid.UUID            `json:"headOfDepartmentId"`
	HeadOfDepartment   *HeadOfDepartmentInfo `json:"headOfDepartment,omitempty"`
	ParentDepartmentID *uuid.UUID            `json:"parentDepartmentId"`
	BudgetCode         string                `json:"budgetCode"`
	TotalEmployee      int                   `json:"totalEmployee"`
}

func ToDepartmentResponse(department *models.Department) *DepartmentResponse {
	if department == nil {
		return nil
	}

	response := &DepartmentResponse{
		ID:                 department.ID,
		Code:               department.Code,
		Name:               department.Name,
		HeadOfDepartmentID: department.HeadOfDepartmentID,
		ParentDepartmentID: department.ParentDepartmentID,
		BudgetCode:         department.BudgetCode,
		TotalEmployee:      len(department.Employees),
	}

	// Include head of department info if available
	if department.HeadOfDepartment != nil {
		response.HeadOfDepartment = &HeadOfDepartmentInfo{
			ID:       department.HeadOfDepartment.ID.String(),
			FullName: department.HeadOfDepartment.FullName,
		}
	}

	return response
}
