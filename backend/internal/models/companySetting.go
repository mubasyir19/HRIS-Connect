package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Untuk JSONB array hari kerja
type WorkingDays []string

func (w WorkingDays) Value() (driver.Value, error) {
	return json.Marshal(w)
}

func (w *WorkingDays) Scan(value interface{}) error {
	if value == nil {
		*w = WorkingDays{}
		return nil
	}
	return json.Unmarshal(value.([]byte), w)
}

type CompanySetting struct {
	ID                      uuid.UUID   `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	CompanyName             string      `gorm:"type:varchar(200)" json:"companyName"`
	CompanyLogoURL          string      `gorm:"type:varchar(500)" json:"companyLogoUrl"`
	WorkingHoursStart       string      `gorm:"type:time" json:"workingHoursStart"` // Simpan sebagai string time
	WorkingHoursEnd         string      `gorm:"type:time" json:"workingHoursEnd"`
	WorkingDays             WorkingDays `gorm:"type:jsonb" json:"workingDays"`
	DefaultLeaveQuota       int         `gorm:"type:int;default:12" json:"defaultLeaveQuota"`
	MaxConsecutiveLeaveDays int         `gorm:"type:int;default:30" json:"maxConsecutiveLeaveDays"`
	Timezone                string      `gorm:"type:varchar(50);default:Asia/Jakarta" json:"timezone"`
	FiscalYearStart         *time.Time  `gorm:"type:date" json:"fiscalYearStart"`
	CreatedAt               time.Time   `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt               time.Time   `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (c *CompanySetting) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}

	if c.WorkingHoursStart == "" {
		c.WorkingHoursStart = "09:00:00"
	}

	if c.WorkingHoursEnd == "" {
		c.WorkingHoursEnd = "17:00:00"
	}

	if len(c.WorkingDays) == 0 {
		c.WorkingDays = WorkingDays{"monday", "tuesday", "wednesday", "thursday", "friday"}
	}

	return nil
}

func (c *CompanySetting) TableName() string {
	return "company_settings"
}
