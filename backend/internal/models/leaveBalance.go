package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LeaveBalance struct {
	ID                 uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	EmployeeID         uuid.UUID      `gorm:"type:uuid;not null;uniqueIndex:idx_employee_year" json:"employeeId"`
	Year               int            `gorm:"type:int;not null;uniqueIndex:idx_employee_year" json:"year"`
	AnnualLeaveQuota   int            `gorm:"type:int;default:12" json:"annualLeaveQuota"`
	AnnualLeaveUsed    int            `gorm:"type:int;default:0" json:"annualLeaveUsed"`
	SickLeaveQuota     int            `gorm:"type:int;default:14" json:"sickLeaveQuota"`
	SickLeaveUsed      int            `gorm:"type:int;default:0" json:"sickLeaveUsed"`
	UnpaidLeaveUsed    int            `gorm:"type:int;default:0" json:"unpaidLeaveUsed"`
	RemainingCarryOver int            `gorm:"type:int;default:0" json:"remainingCarryOver"`
	CreatedAt          time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationship
	Employee *Employee `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
}

func (l *LeaveBalance) BeforeCreate(tx *gorm.DB) error {
	if l.ID == uuid.Nil {
		l.ID = uuid.New()
	}
	return nil
}

func (l *LeaveBalance) TableName() string {
	return "leave_balances"
}
