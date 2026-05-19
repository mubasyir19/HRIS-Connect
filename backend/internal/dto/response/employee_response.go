package response

import (
	"backend/internal/models"
	"time"
)

type EmployeeResponse struct {
	// Personal Information
	ID            string     `json:"id"`
	NIK           string     `json:"nik"`
	FullName      string     `json:"fullname"`
	Email         string     `json:"email"`
	Phone         string     `json:"phone"`
	Address       string     `json:"address"`
	BirthDate     *time.Time `json:"birthDate"`
	BirthPlace    string     `json:"birthPlace"`
	Gender        string     `json:"gender"`
	MaritalStatus string     `json:"maritalStatus"`

	// Employment Information
	Department       string     `json:"department"`
	Position         string     `json:"position"`
	JobLevel         string     `json:"jobLevel"`
	ManagerID        *string    `json:"managerId"`
	ManagerName      string     `json:"managerName,omitempty"`
	JoinDate         time.Time  `json:"joinDate"`
	ContractEndDate  *time.Time `json:"contractEndDate"`
	EmploymentStatus string     `json:"employmentStatus"`

	// Leave & Attendance
	RemainingLeaveDays int `json:"remainingLeaveDays"`
	RemainingSickDays  int `json:"remainingSickDays"`

	// Emergency Contact
	EmergencyContactName  string `json:"emergencyContactName"`
	EmergencyContactPhone string `json:"emergencyContactPhone"`

	// Bank Information
	BankName          string `json:"bankName"`
	BankAccountNumber string `json:"bankAccountNumber"`
	BankAccountName   string `json:"bankAccountName"`

	// Metadata
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

func ToDetailEmployeeResponse(employee *models.Employee) *EmployeeResponse {
	if employee == nil {
		return nil
	}

	response := &EmployeeResponse{
		ID:            employee.ID.String(),
		NIK:           employee.NIK,
		FullName:      employee.FullName,
		Email:         employee.Email,
		Phone:         employee.Phone,
		Address:       employee.Address,
		BirthDate:     employee.BirthDate,
		BirthPlace:    employee.BirthPlace,
		Gender:        string(employee.Gender),
		MaritalStatus: string(employee.MaritalStatus),

		// Employment Information
		Department:       employee.DepartmentID.String(),
		Position:         employee.Position,
		JobLevel:         employee.JobLevel,
		JoinDate:         employee.JoinDate,
		ContractEndDate:  employee.ContractEndDate,
		EmploymentStatus: string(employee.EmploymentStatus),

		// Leave & Attendance
		RemainingLeaveDays: employee.RemainingLeaveDays,
		RemainingSickDays:  employee.RemainingSickDays,

		// Emergency Contact
		EmergencyContactName:  employee.EmergencyContactName,
		EmergencyContactPhone: employee.EmergencyContactPhone,

		// Bank Information
		BankName:          employee.BankName,
		BankAccountNumber: employee.BankAccountNumber,
		BankAccountName:   employee.BankAccountName,

		// Metadata
		CreatedAt: employee.CreatedAt,
		UpdatedAt: employee.UpdatedAt,
	}

	return response
}

func ToListEmployeeResponse(employees []models.Employee) []EmployeeResponse {
	if len(employees) == 0 {
		return []EmployeeResponse{}
	}

	response := make([]EmployeeResponse, 0, len(employees))
	for _, employee := range employees {
		response = append(response, *ToDetailEmployeeResponse(&employee))
	}

	return response
}
