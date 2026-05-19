package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GenderEmployee string

const (
	GenderMale   GenderEmployee = "Male"
	GenderFemale GenderEmployee = "Female"
)

type MaritalStatusEmployee string

const (
	MaritalSingle   MaritalStatusEmployee = "Single"
	MaritalMarried  MaritalStatusEmployee = "Married"
	MaritalDivorced MaritalStatusEmployee = "Divorced"
	MaritalWodowed  MaritalStatusEmployee = "Widowed"
)

type EmploymentStatus string

const (
	StatusPermanent  EmploymentStatus = "Permanent"
	StatusContract   EmploymentStatus = "Contract"
	StatusProbation  EmploymentStatus = "Probation"
	StatusInternship EmploymentStatus = "Internship"
)

type Employee struct {
	ID            uuid.UUID             `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	NIK           string                `gorm:"type:varchar(50);unique;not null" json:"nik"`
	FullName      string                `gorm:"type:varchar(100);not null" json:"fullname"`
	Email         string                `gorm:"type:varchar(100);unique;not null" json:"email"`
	Phone         string                `gorm:"type:varchar(20)" json:"phone"`
	Address       string                `gorm:"type:text" json:"address"`
	BirthDate     *time.Time            `gorm:"type:date" json:"birthDate"`
	BirthPlace    string                `gorm:"type:varchar(50)" json:"birthPlace"`
	Gender        GenderEmployee        `gorm:"type:varchar(10);not null;default:Male" json:"gender"`
	MaritalStatus MaritalStatusEmployee `gorm:"type:varchar(20);not null;default:Single" json:"maritalStatus"`

	// Employee Information
	// Department       string           `gorm:"varchar(50);not null" json:"department"`
	DepartmentID     *uuid.UUID       `gorm:"type:uuid" json:"departmentId"`
	Department       *Department      `gorm:"foreignKey:DepartmentID" json:"department,omitempty"`
	Position         string           `gorm:"varchar(100);not null" json:"position"`
	JobLevel         string           `gorm:"varchar(50)" json:"jobLevel"`
	ManagerID        *uuid.UUID       `gorm:"type:uuid" json:"managerId"`
	JoinDate         time.Time        `gorm:"type:date;not null" json:"joinDate"`
	ContractEndDate  *time.Time       `gorm:"type:date" json:"contractEndDate"`
	EmploymentStatus EmploymentStatus `gorm:"type:varchar(20);not null;default:Contract" json:"employmentStatus"`

	// Leave & Attendance
	RemainingLeaveDays int `gorm:"type:integer;default:12" json:"remainingLeaveDays"`
	RemainingSickDays  int `gorm:"type:integer;default:14" json:"remainingSickDays"`

	// Emergency Contact
	EmergencyContactName  string `gorm:"type:varchar(100)" json:"emergencyContactName"`
	EmergencyContactPhone string `gorm:"type:varchar(20)" json:"emergencyContactPhone"`

	// Bank Information
	BankName          string `gorm:"type:varchar(50)" json:"bankName"`
	BankAccountNumber string `gorm:"type:varchar(50)" json:"bankAccountNumber"`
	BankAccountName   string `gorm:"type:varchar(100)" json:"bankAccountName"`

	// Metadata
	CreatedAt time.Time      `gorm:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Manager       *Employee      `gorm:"foreignKey:ManagerID" json:"manager,omitempty"`
	User          *User          `gorm:"foreignKey:EmployeeID" json:"user,omitempty"`
	LeaveRequests []LeaveRequest `gorm:"foreignKey:EmployeeID" json:"leaveRequests,omitempty"`
	LeaveBalances []LeaveBalance `gorm:"foreignKey:EmployeeID" json:"leaveBalances,omitempty"`
	Attendances   []Attendance   `gorm:"foreignKey:EmployeeID" json:"attendances,omitempty"`
}

func (e *Employee) BeforeCreate(tx *gorm.DB) error {
	if e.ID == uuid.Nil {
		e.ID = uuid.New()
	}
	if e.Gender == "" {
		e.Gender = GenderMale
	}
	if e.MaritalStatus == "" {
		e.MaritalStatus = MaritalSingle
	}
	if e.EmploymentStatus == "" {
		e.EmploymentStatus = StatusContract
	}
	return nil
}

func (e *Employee) TableName() string {
	return "employees"
}
