package requests

import (
	"backend/internal/models"
	"time"

	"github.com/google/uuid"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UpdateUserRequest struct {
	ID           uuid.UUID   `json:"id"`
	EmployeeID   *uuid.UUID  `json:"employeeId"`
	Email        *string     `json:"email"`
	Password     *string     `json:"-"`
	Role         models.Role `json:"role"`
	IsActive     *bool       `json:"isActive"`
	LastLogin    *time.Time  `json:"lastLogin"`
	RefreshToken *string     `json:"refreshToken"`
}
