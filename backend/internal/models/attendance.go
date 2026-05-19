package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AttendanceStatus string

const (
	AttendancePresent AttendanceStatus = "present"
	AttendanceAbsent  AttendanceStatus = "absent"
	AttendanceLate    AttendanceStatus = "late"
	AttendanceHalfDay AttendanceStatus = "half_day"
	AttendanceHoliday AttendanceStatus = "holiday"
	AttendanceSick    AttendanceStatus = "sick"
	AttendanceLeave   AttendanceStatus = "leave"
)

type Attendance struct {
	ID               uuid.UUID        `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	EmployeeID       uuid.UUID        `gorm:"type:uuid;not null;uniqueIndex:idx_employee_date" json:"employeeId"`
	Date             time.Time        `gorm:"type:date;not null;uniqueIndex:idx_employee_date" json:"date"`
	CheckInTime      *time.Time       `gorm:"type:time" json:"checkInTime"`
	CheckOutTime     *time.Time       `gorm:"type:time" json:"checkOutTime"`
	CheckInLocation  string           `gorm:"type:varchar(100)" json:"checkInLocation"` // Simpan sebagai string POINT
	CheckOutLocation string           `gorm:"type:varchar(100)" json:"checkOutLocation"`
	WorkHours        float64          `gorm:"type:decimal(5,2)" json:"workHours"`
	OvertimeHours    float64          `gorm:"type:decimal(5,2)" json:"overtimeHours"`
	Status           AttendanceStatus `gorm:"type:varchar(20);not null;default:present;index" json:"status"`
	Notes            string           `gorm:"type:text" json:"notes"`
	CreatedAt        time.Time        `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt        time.Time        `gorm:"autoUpdateTime" json:"updatedAt"`

	// Relationship
	Employee *Employee `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
}

func (a *Attendance) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	if a.Status == "" {
		a.Status = AttendancePresent
	}
	return nil
}

func (a *Attendance) TableName() string {
	return "attendances"
}
