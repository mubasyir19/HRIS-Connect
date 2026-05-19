package services

import (
	"backend/internal/dto/requests"
	"backend/internal/models"
	"backend/internal/repositories"
	"errors"

	"github.com/google/uuid"
)

type DepartmentService interface {
	GetAll(filter map[string]interface{}, page, limit int) ([]models.Department, int64, error)
	GetByID(id uuid.UUID) (*models.Department, error)
	AddNewDepartment(req *requests.CreateNewDepartment) (*models.Department, error)
	Update(id uuid.UUID, req requests.UpdateDepartment) (*models.Department, error)
	Delete(id uuid.UUID)
}

type departmentService struct {
	repository   repositories.DepartmentRepository
	employeeRepo repositories.EmployeeRepository
}

func NewDepartmentService(repository repositories.DepartmentRepository, employeeRepo repositories.EmployeeRepository) DepartmentService {
	return &departmentService{repository, employeeRepo}
}

func (s *departmentService) GetAll(filter map[string]interface{}, page, limit int) ([]models.Department, int64, error) {
	// validate parameter
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	departments, total, err := s.repository.GetAll(filter, page, limit)
	if err != nil {
		return nil, 0, errors.New("failed to get department")
	}

	return departments, total, nil
}

func (s *departmentService) GetByID(id uuid.UUID) (*models.Department, error) {
	if id != uuid.Nil {
		return nil, errors.New("")
	}

	department, err := s.repository.GetByID(id.String())
	if err != nil {
		return nil, errors.New("data doesn't exist")
	}

	return department, nil
}

func (s *departmentService) AddNewDepartment(req *requests.CreateNewDepartment) (*models.Department, error) {
	// validate request
	if req.Code == "" {
		return nil, errors.New("code is required")
	}

	if req.Name == "" {
		return nil, errors.New("name is required")
	}

	// check duplicate with code
	_, err := s.repository.GetByCode(req.Code)
	if err != nil {
		return nil, errors.New("this department already exist")
	}

	// validate parent department id
	if req.ParentDepartmentID != nil {
		_, err := s.repository.GetByID(req.ParentDepartmentID.String())
		if err != nil {
			return nil, errors.New("parent department doesn't exist")
		}
	}

	// validate head of department id
	if req.HeadOfDepartmentID != nil {
		_, err := s.repository.GetByID(req.HeadOfDepartmentID.String())
		if err != nil {
			return nil, errors.New("head of department doesn't exist")
		}
	}

	// payload add new
	newDepartment := &models.Department{
		Code:               req.Code,
		Name:               req.Name,
		ParentDepartmentID: req.ParentDepartmentID,
		HeadOfDepartmentID: req.ParentDepartmentID,
		BudgetCode:         req.BudgetCode,
	}

	// create department
	if err := s.repository.Create(newDepartment); err != nil {
		return nil, errors.New("failed to add new department")
	}

	// return
	return newDepartment, nil
}

func (s *departmentService) Update(id uuid.UUID, req requests.UpdateDepartment) (*models.Department, error) {
	if id != uuid.Nil {
		return nil, errors.New("id is required")
	}

	// check duplication with code
	checkDeparment, err := s.repository.GetByID(id.String())
	if err != nil {
		return nil, errors.New("data not found")
	}

	// update data
	if req.Code != nil {
		checkDeparment.Code = *req.Code
	}

	if req.Name != nil {
		checkDeparment.Name = *req.Name
	}

	if req.ParentDepartmentID != nil {
		checkDeparment.ParentDepartmentID = req.ParentDepartmentID
	}

	if req.HeadOfDepartmentID != nil {
		checkDeparment.HeadOfDepartmentID = req.HeadOfDepartmentID
	}

	if err := s.repository.Update(checkDeparment); err != nil {
		return nil, errors.New("failed to update department")
	}

	return checkDeparment, nil
}

func (s *departmentService) Delete(id uuid.UUID) {}
