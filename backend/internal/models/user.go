package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin    Role = "admin"
	RoleManager  Role = "manager"
	RoleEmployee Role = "employee"
)

type User struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	EmployeeID   uuid.UUID      `gorm:"type:uuid;uniqueIndex;not null" json:"employeeId"`
	Email        string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password     string         `gorm:"type:varchar(255);not null" json:"-"`
	Role         Role           `gorm:"type:varchar(20);not null;default:employee" json:"role"`
	IsActive     bool           `gorm:"default:true" json:"isActive"`
	LastLogin    *time.Time     `gorm:"type:timestamp" json:"lastLogin"`
	RefreshToken string         `gorm:"type:varchar(500)" json:"-"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationship
	Employee       *Employee      `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
	AuditLogs      []AuditLog     `gorm:"foreignKey:UserID" json:"-"`
	ApprovedLeaves []LeaveRequest `gorm:"foreignKey:ApprovedBy" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	if u.Role == "" {
		u.Role = RoleEmployee
	}
	return nil
}

func (u *User) TableName() string {
	return "users"
}
