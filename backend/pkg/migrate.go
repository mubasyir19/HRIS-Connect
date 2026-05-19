package pkg

import (
	"backend/internal/models"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.Employee{},
		&models.User{},
		&models.LeaveRequest{},
		&models.LeaveBalance{},
		&models.Attendance{},
		&models.Department{},
		&models.Position{},
		&models.CompanyHoliday{},
		&models.AuditLog{},
		&models.CompanySetting{},
	); err != nil {
		return err
	}

	relationships := []struct {
		model interface{}
		name  string
	}{
		{&models.Department{}, "HeadOfDepartment"},
		{&models.Department{}, "ParentDepartment"},
		{&models.Employee{}, "Department"},
		{&models.Employee{}, "Manager"},
		{&models.User{}, "Employee"},
		{&models.AuditLog{}, "User"},
		{&models.Attendance{}, "Employee"},
		{&models.LeaveRequest{}, "Employee"},
		{&models.LeaveRequest{}, "Approver"},
		{&models.LeaveBalance{}, "Employee"},
		{&models.Position{}, "Department"},
	}

	for _, rel := range relationships {
		if err := db.Migrator().CreateConstraint(rel.model, rel.name); err != nil {
			return err
		}
	}

	return nil
}
