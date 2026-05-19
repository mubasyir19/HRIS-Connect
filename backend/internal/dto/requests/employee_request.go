package requests

import (
	"time"

	"github.com/google/uuid"
)

type CreateEmployeeRequest struct {
	// Personal Information
	NIK           string     `json:"nik" validate:"required,min=3,max=50"`
	FullName      string     `json:"fullname" validate:"required,min=2,max=100"` // required
	Email         string     `json:"email" validate:"required,email"`            // required
	Phone         string     `json:"phone" validate:"omitempty,min=10,max=15"`
	Address       string     `json:"address" validate:"omitempty"`
	BirthDate     *time.Time `json:"birthDate" validate:"omitempty"`
	BirthPlace    string     `json:"birthPlace" validate:"omitempty"`
	Gender        string     `json:"gender" validate:"omitempty,oneof=Male Female"`
	MaritalStatus string     `json:"maritalStatus" validate:"omitempty,oneof=Single Married Divorced Widowed"`

	// Employment Information
	DepartmentID     *uuid.UUID `json:"departmentId" validate:"required"` // required
	Position         string     `json:"position" validate:"required"`     // required
	JobLevel         string     `json:"jobLevel" validate:"omitempty"`
	ManagerID        *string    `json:"managerId" validate:"omitempty"`
	JoinDate         time.Time  `json:"joinDate" validate:"required"` // required
	ContractEndDate  *time.Time `json:"contractEndDate" validate:"omitempty"`
	EmploymentStatus string     `json:"employmentStatus" validate:"required,oneof=Permanent Contract Probation Internship"`

	// Leave & Attendance (initial balance)
	RemainingLeaveDays int `json:"remainingLeaveDays" validate:"omitempty,min=0"`
	RemainingSickDays  int `json:"remainingSickDays" validate:"omitempty,min=0"`

	// Emergency Contact
	EmergencyContactName  string `json:"emergencyContactName" validate:"omitempty"`
	EmergencyContactPhone string `json:"emergencyContactPhone" validate:"omitempty"`

	// Bank Information
	BankName          string `json:"bankName" validate:"omitempty"`
	BankAccountNumber string `json:"bankAccountNumber" validate:"omitempty"`
	BankAccountName   string `json:"bankAccountName" validate:"omitempty"`

	// User Account (opsional, untuk auto-create user)
	CreateUserAccount bool   `json:"createUserAccount" validate:"omitempty"`
	Password          string `json:"password" validate:"omitempty,min=8"`
	Role              string `json:"role" validate:"omitempty,oneof=admin manager employee"`
}

type UpdateEmployeeRequest struct {
}
