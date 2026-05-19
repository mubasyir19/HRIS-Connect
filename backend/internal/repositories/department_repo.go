package repositories

import (
	"backend/internal/models"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DepartmentRepository interface {
	GetAll(filter map[string]interface{}, page, limit int) ([]models.Department, int64, error)
	GetByID(id string) (*models.Department, error)
	GetByCode(code string) (*models.Department, error)
	Create(depatment *models.Department) error
	Update(depatment *models.Department) error
	Delete(id uuid.UUID) error
}

type departmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentRepository{db}
}

func (r *departmentRepository) GetAll(filter map[string]interface{}, page, limit int) ([]models.Department, int64, error) {
	var departments []models.Department
	var total int64

	count := r.db.Model(&models.Department{}).Where(filter).Count(&total)
	if count.Error != nil {
		return nil, 0, fmt.Errorf("failed to count departments: %w", count.Error)
	}

	result := r.db.Where(filter).Offset((page - 1) * limit).Limit(limit).Find(&departments)
	if result.Error != nil {
		return nil, 0, fmt.Errorf("result get departmnets: %w", result.Error)
	}

	return departments, 0, nil
}

func (r *departmentRepository) GetByID(id string) (*models.Department, error) {
	var department models.Department
	result := r.db.Where("id = ? ", id).First(&department)
	if result.Error != nil {
		return nil, fmt.Errorf("get by id department: %w", result.Error)
	}

	return &department, nil
}

func (r *departmentRepository) GetByCode(code string) (*models.Department, error) {
	var department models.Department
	result := r.db.Where("code = ? ", code).First(&department)
	if result.Error != nil {
		return nil, fmt.Errorf("get by code department: %w", result.Error)
	}

	return &department, nil
}

func (r *departmentRepository) Create(department *models.Department) error {
	result := r.db.Create(department)
	if result.Error != nil {
		return fmt.Errorf("failed to create department: %w", result.Error)
	}

	return nil
}

func (r *departmentRepository) Update(department *models.Department) error {
	result := r.db.Save(&department)
	if result.Error != nil {
		return fmt.Errorf("failed get update department: %w", result.Error)
	}

	return nil
}

func (r *departmentRepository) Delete(id uuid.UUID) error {
	result := r.db.Where("id = ?", id).Delete(&models.Department{})
	if result.Error != nil {
		return fmt.Errorf("failed delete deparmetnt: %w", result.Error)
	}

	return nil
}
