package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HolidayType string

const (
	HolidayNational  HolidayType = "national"
	HolidayCompany   HolidayType = "company"
	HolidayReligious HolidayType = "religious"
)

type CompanyHoliday struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name        string         `gorm:"type:varchar(100);not null" json:"name"`
	Date        time.Time      `gorm:"type:date;not null;uniqueIndex:idx_date_name" json:"date"`
	Type        HolidayType    `gorm:"type:varchar(20);not null" json:"type"`
	IsRecurring bool           `gorm:"default:false" json:"isRecurring"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (h *CompanyHoliday) BeforeCreate(tx *gorm.DB) error {
	if h.ID == uuid.Nil {
		h.ID = uuid.New()
	}
	return nil
}

func (h *CompanyHoliday) TableName() string {
	return "company_holidays"
}
