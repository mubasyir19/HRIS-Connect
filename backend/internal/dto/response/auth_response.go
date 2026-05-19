package response

import (
	"backend/internal/models"

	"github.com/google/uuid"
)

type DataUserResponse struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"fullname"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
}

type LoginResponse struct {
	AccessToken  string           `json:"accessToken"`
	RefreshToken string           `json:"refreshToken"`
	Data         DataUserResponse `json:"data"`
}

func ToAuthResponse(user *models.User, employee *models.Employee) *DataUserResponse {
	if user == nil {
		return nil
	}

	fullName := ""
	if employee != nil {
		fullName = employee.FullName
	}

	return &DataUserResponse{
		ID:       user.ID,
		FullName: fullName,
		Email:    user.Email,
		Role:     string(user.Role),
	}
}
