package response

import "backend/internal/models"

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func ToUserResponse(user *models.User) *UserResponse {
	return &UserResponse{
		ID:    user.ID.String(),
		Email: user.Email,
		Role:  string(user.Role),
	}
}
