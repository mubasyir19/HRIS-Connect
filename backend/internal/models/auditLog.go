package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuditLog struct {
	ID         uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID     *uuid.UUID      `gorm:"type:uuid;index" json:"userId"`
	Action     string          `gorm:"type:varchar(100);not null;index" json:"action"`
	EntityType string          `gorm:"type:varchar(50);not null;index" json:"entityType"`
	EntityID   *uuid.UUID      `gorm:"type:uuid;index" json:"entityId"`
	OldData    json.RawMessage `gorm:"type:jsonb" json:"oldData,omitempty"`
	NewData    json.RawMessage `gorm:"type:jsonb" json:"newData,omitempty"`
	IPAddress  string          `gorm:"type:inet" json:"ipAddress"`
	UserAgent  string          `gorm:"type:text" json:"userAgent"`
	CreatedAt  time.Time       `gorm:"autoCreateTime;index" json:"createdAt"`

	// Relationship
	User *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (a *AuditLog) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return nil
}

func (a *AuditLog) TableName() string {
	return "audit_logs"
}
