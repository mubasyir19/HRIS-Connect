package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Department struct {
	ID                 uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Code               string         `gorm:"type:varchar(10);uniqueIndex;not null" json:"code"`
	Name               string         `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"`
	HeadOfDepartmentID *uuid.UUID     `gorm:"type:uuid" json:"headOfDepartmentId"`
	ParentDepartmentID *uuid.UUID     `gorm:"type:uuid" json:"parentDepartmentId"`
	BudgetCode         string         `gorm:"type:varchar(50)" json:"budgetCode"`
	CreatedAt          time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	HeadOfDepartment *Employee    `gorm:"foreignKey:HeadOfDepartmentID" json:"headOfDepartment,omitempty"`
	ParentDepartment *Department  `gorm:"foreignKey:ParentDepartmentID" json:"parentDepartment,omitempty"`
	SubDepartments   []Department `gorm:"foreignKey:ParentDepartmentID" json:"subDepartments,omitempty"`
	Employees        []Employee   `gorm:"foreignKey:DepartmentID" json:"employees,omitempty"`
}

func (d *Department) BeforeCreate(tx *gorm.DB) error {
	if d.ID == uuid.Nil {
		d.ID = uuid.New()
	}
	return nil
}

func (d *Department) TableName() string {
	return "departments"
}
