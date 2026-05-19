package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Position struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Code         string         `gorm:"type:varchar(10);uniqueIndex;not null" json:"code"`
	Name         string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"name"`
	JobLevel     int            `gorm:"type:int;check:job_level BETWEEN 1 AND 15" json:"jobLevel"`
	DepartmentID *uuid.UUID     `gorm:"type:uuid" json:"departmentId"`
	MinSalary    float64        `gorm:"type:decimal(15,2)" json:"minSalary"`
	MaxSalary    float64        `gorm:"type:decimal(15,2)" json:"maxSalary"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Department *Department `gorm:"foreignKey:DepartmentID" json:"department,omitempty"`
}

func (p *Position) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}

func (p *Position) TableName() string {
	return "positions"
}
