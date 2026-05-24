package services

import (
	"backend/internal/dto/requests"
	"backend/internal/dto/response"
	"backend/internal/models"
	"backend/internal/repositories"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type EmployeeService interface {
	GetAll(filter map[string]interface{}, page, limit int) ([]models.Employee, int64, error)
	Create(req requests.CreateEmployeeRequest) (*response.EmployeeResponse, error)                // output return pointer response employee + error
	GetByID(id uuid.UUID) (*response.EmployeeResponse, error)                                     // output return pointer response employee + error
	GetByNIK(nik string) (*response.EmployeeResponse, error)                                      // output return pointer response employee + error
	Update(id uuid.UUID, req *requests.UpdateEmployeeRequest) (*response.EmployeeResponse, error) // output return pointer response employee + error
	Delete(id uuid.UUID) error                                                                    // output error
	UpdateLeaveBalance()                                                                          // output = error
}

type employeeService struct {
	employeeRepo repositories.EmployeeRepository
	userRepo     repositories.UserRepository
	db           *gorm.DB
}

func NewEmployeeService(employeeRepo repositories.EmployeeRepository, userRepo repositories.UserRepository, db *gorm.DB) EmployeeService {
	return &employeeService{
		employeeRepo: employeeRepo,
		userRepo:     userRepo,
		db:           db,
	}
}

func (s *employeeService) GetAll(filter map[string]interface{}, page, limit int) ([]models.Employee, int64, error) {
	// validate parameter
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	employees, total, err := s.employeeRepo.GetAll(filter, page, limit)
	if err != nil {
		return nil, 0, errors.New("failed get list employees")
	}

	return employees, total, nil
}

func (s *employeeService) Create(req requests.CreateEmployeeRequest) (*response.EmployeeResponse, error) {
	// Cek field request tidak boleh kosong yang required

	// cek employee by id atau email
	existingEmployee, _ := s.employeeRepo.GetByEmail(req.Email)
	// jika sudah terdaftar, maka return already exist
	if existingEmployee != nil {
		return nil, errors.New("Employee already exist")
	}

	// cek validasi request yang required
	if req.FullName == "" {
		return nil, errors.New("fullname employee is required")
	}

	if req.Email == "" {
		return nil, errors.New("email employee is required")
	}

	// if req.Password == "" {
	// 	return nil, errors.New("password is required")
	// }

	if req.DepartmentID == nil {
		return nil, errors.New("departmentId is required")
	}

	if req.Position == "" {
		return nil, errors.New("position employee is required")
	}

	if req.JoinDate.IsZero() {
		return nil, errors.New("join date employee is required")
	}

	// buat payload untuk menampung semua field
	employee := &models.Employee{
		ID:           uuid.New(),
		NIK:          req.NIK,
		FullName:     req.FullName,
		Email:        req.Email,
		DepartmentID: req.DepartmentID,
		Position:     req.Position,
		JoinDate:     req.JoinDate,
	}

	// Transaction = membuat 2 data => Employee & User
	err := s.db.Transaction(func(tx *gorm.DB) error {
		// create employee
		if err := s.employeeRepo.Create(employee); err != nil {
			return errors.New("failed create new employee")
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return errors.New("failed to hash Password")
		}

		user := &models.User{
			ID:         uuid.New(),
			EmployeeID: employee.ID,
			Email:      employee.Email,
			Password:   string(hashedPassword),
			Role:       models.Role(req.Role),
			IsActive:   true,
		}

		// create user
		if err := s.userRepo.Create(user); err != nil {
			return errors.New("failed create new user")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return response.ToDetailEmployeeResponse(employee), nil
}

func (s *employeeService) GetByID(id uuid.UUID) (*response.EmployeeResponse, error) {
	employee, err := s.employeeRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return response.ToDetailEmployeeResponse(employee), nil
}

func (s *employeeService) GetByNIK(nik string) (*response.EmployeeResponse, error) {
	employee, err := s.employeeRepo.GetByNIK(nik)
	if err != nil {
		return nil, err
	}

	return response.ToDetailEmployeeResponse(employee), nil
}

func (s *employeeService) Update(id uuid.UUID, req *requests.UpdateEmployeeRequest) (*response.EmployeeResponse, error) {
	// cek employee ada atau tidak dengan GetByID
	// return jika errors.New() jika tidak ada
	existingEmployee, err := s.employeeRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("employee not found")
	}

	// validate kosong atau tidak untuk field request nya yang required
	if existingEmployee.FullName == "" {
		return nil, errors.New("fullname employee is required")
	}

	if existingEmployee.Email == "" {
		return nil, errors.New("email employee is required")
	}

	if existingEmployee.DepartmentID == nil {
		return nil, errors.New("departmentId is required")
	}

	if existingEmployee.Position == "" {
		return nil, errors.New("position employee is required")
	}

	if existingEmployee.JoinDate.IsZero() {
		return nil, errors.New("join date employee is required")
	}

	// update dengan panggil repo Update
	if err := s.employeeRepo.Update(existingEmployee); err != nil {
		return nil, errors.New("failed update employee")
	}

	// return response, nil
	return response.ToDetailEmployeeResponse(existingEmployee), nil
}

func (s *employeeService) Delete(id uuid.UUID) error {
	// cek employee ada atau tidak
	existingEmployee, err := s.employeeRepo.GetByID(id)
	if err != nil {
		return errors.New("employee not found")
	}

	// repo delete employee
	if err := s.employeeRepo.Delete(existingEmployee.ID); err != nil {
		return errors.New("failed delete employee")
	}

	return nil
}

func (s *employeeService) UpdateLeaveBalance() {}
