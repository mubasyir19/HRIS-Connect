package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LeaveType string

const (
	LeaveAnnual    LeaveType = "annual"
	LeaveSick      LeaveType = "sick"
	LeaveUnpaid    LeaveType = "unpaid"
	LeaveMaternity LeaveType = "maternity"
	LeavePaternity LeaveType = "paternity"
	LeaveEmergency LeaveType = "emergency"
	LeaveStudy     LeaveType = "study"
)

type LeaveStatus string

const (
	LeavePending   LeaveStatus = "pending"
	LeaveApproved  LeaveStatus = "approved"
	LeaveRejected  LeaveStatus = "rejected"
	LeaveCancelled LeaveStatus = "cancelled"
)

type LeaveRequest struct {
	ID              uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	EmployeeID      uuid.UUID      `gorm:"type:uuid;not null;index" json:"employeeId"`
	LeaveType       LeaveType      `gorm:"type:varchar(30);not null" json:"leaveType"`
	StartDate       time.Time      `gorm:"type:date;not null;index" json:"startDate"`
	EndDate         time.Time      `gorm:"type:date;not null" json:"endDate"`
	TotalDays       int            `gorm:"type:int;not null" json:"totalDays"`
	Reason          string         `gorm:"type:text" json:"reason"`
	AttachmentURL   string         `gorm:"type:varchar(500)" json:"attachmentUrl"`
	Status          LeaveStatus    `gorm:"type:varchar(20);not null;default:pending;index" json:"status"`
	ApprovedBy      *uuid.UUID     `gorm:"type:uuid" json:"approvedBy"`
	ApprovedAt      *time.Time     `gorm:"type:timestamp" json:"approvedAt"`
	RejectionReason string         `gorm:"type:text" json:"rejectionReason"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Employee *Employee `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
	Approver *User     `gorm:"foreignKey:ApprovedBy" json:"approver,omitempty"`
}

func (l *LeaveRequest) BeforeCreate(tx *gorm.DB) error {
	if l.ID == uuid.Nil {
		l.ID = uuid.New()
	}
	if l.Status == "" {
		l.Status = LeavePending
	}
	return nil
}

func (l *LeaveRequest) TableName() string {
	return "leave_requests"
}
