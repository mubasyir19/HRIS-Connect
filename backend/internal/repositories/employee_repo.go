package repositories

import (
	"backend/internal/models"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(employee *models.Employee) error
	GetByID(id uuid.UUID) (*models.Employee, error)
	GetByEmail(email string) (*models.Employee, error)
	GetByNIK(nik string) (*models.Employee, error)
	GetAll(filter map[string]interface{}, page, limit int) ([]models.Employee, int64, error)
	Update(employee *models.Employee) error
	Delete(id uuid.UUID) error
	GetByDepartment(department string) ([]models.Employee, error)
	GetByManagerID(managerID uuid.UUID) ([]models.Employee, error)
	UpdateLeaveBalance(id uuid.UUID, days int) error
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db}
}

func (r *employeeRepository) Create(employee *models.Employee) error {
	return r.db.Create(employee).Error
}

func (r *employeeRepository) GetByID(id uuid.UUID) (*models.Employee, error) {
	var employee models.Employee
	// if err := r.db.Where("id = ?", id).First(&employee).Error; err != nil {
	if err := r.db.Preload("User").Where("id = ?", id).First(&employee).Error; err != nil {
		return nil, fmt.Errorf("get by id employee : %w", err)
	}

	return &employee, nil
}

func (r *employeeRepository) GetByEmail(email string) (*models.Employee, error) {
	var employee models.Employee
	if err := r.db.Preload("User").Where("email = ?", email).First(&employee).Error; err != nil {
		return nil, fmt.Errorf("get by email employee : %w", err)
	}

	return &employee, nil
}

func (r *employeeRepository) GetByNIK(nik string) (*models.Employee, error) {
	var employee models.Employee
	if err := r.db.Where("nik = ?", nik).First(&employee).Error; err != nil {
		return nil, fmt.Errorf("get by nik employee : %w", err)
	}

	return &employee, nil
}

func (r *employeeRepository) GetAll(filter map[string]interface{}, page, limit int) ([]models.Employee, int64, error) {
	var employees []models.Employee
	var total int64

	if err := r.db.Model(&models.Employee{}).Where(filter).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count employees: %w", err)
	}

	if err := r.db.Where(filter).Offset((page - 1) * limit).Limit(limit).Find(&employees).Error; err != nil {
		return nil, 0, fmt.Errorf("get all employees: %w", err)
	}

	return employees, total, nil

}

func (r *employeeRepository) Update(employee *models.Employee) error {
	if err := r.db.Save(&employee).Error; err != nil {
		return fmt.Errorf("update employee: %w", err)
	}

	return nil
}

func (r *employeeRepository) Delete(id uuid.UUID) error {
	if err := r.db.Where("id = ?", id).Delete(&models.Employee{}).Error; err != nil {
		return fmt.Errorf("delete employee: %w", err)
	}

	return nil
}

func (r *employeeRepository) GetByDepartment(department string) ([]models.Employee, error) {
	var employees []models.Employee
	if err := r.db.Where("department = ?", department).Find(&employees).Error; err != nil {
		return nil, fmt.Errorf("get employees by department = %w", err)
	}

	return employees, nil
}

func (r *employeeRepository) GetByManagerID(managerID uuid.UUID) ([]models.Employee, error) {
	var employees []models.Employee
	if err := r.db.Where("managerId = ? ", managerID).Find(&employees).Error; err != nil {
		return nil, fmt.Errorf("get by managerID: %w", err)
	}

	return employees, nil
}

func (r *employeeRepository) UpdateLeaveBalance(id uuid.UUID, days int) error {
	if err := r.db.Where("id = ?", id).Update("remainingLeaveDays", gorm.Expr("remainingLeaveDays + ?", days)).Error; err != nil {
		return fmt.Errorf("update leave balance = %w", err)
	}

	return nil
}
