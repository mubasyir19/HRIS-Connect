package requests

import "github.com/google/uuid"

type CreateNewDepartment struct {
	Code               string     `json:"code"`
	Name               string     `json:"name"`
	HeadOfDepartmentID *uuid.UUID `json:"headOfDepartmentId"`
	ParentDepartmentID *uuid.UUID `json:"parentDepartmentId"`
	BudgetCode         string     `json:"budgetCode"`
}

type UpdateDepartment struct {
	Code               *string    `json:"code"`
	Name               *string    `json:"name"`
	HeadOfDepartmentID *uuid.UUID `json:"headOfDepartmentId"`
	ParentDepartmentID *uuid.UUID `json:"parentDepartmentId"`
	BudgetCode         *string    `json:"budgetCode"`
}
