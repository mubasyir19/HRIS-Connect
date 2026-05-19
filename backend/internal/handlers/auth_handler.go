package handlers

import (
	"backend/internal/dto/requests"
	"backend/internal/middleware"
	"backend/internal/models"
	"backend/internal/services"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type userHandler struct {
	authService     services.AuthService
	employeeService services.EmployeeService
}

func NewAuthHandler(authService services.AuthService, employeeService services.EmployeeService) *userHandler {
	return &userHandler{authService, employeeService}
}

func (h *userHandler) Login(c fiber.Ctx) error {
	var req requests.LoginRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request format",
			"details": err.Error(),
		})
	}

	loginResponse, err := h.authService.Login(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	isProduction := os.Getenv("ENV") == "production"

	accessCookie := &fiber.Cookie{
		Name:     "access_token",
		Value:    loginResponse.AccessToken,
		Path:     "/",
		HTTPOnly: true,
		Secure:   isProduction,
		SameSite: "None",
		MaxAge:   7 * 24 * 60 * 60,
	}
	c.Cookie(accessCookie)

	refreshCookie := &fiber.Cookie{
		Name:     "refresh_token",
		Value:    loginResponse.RefreshToken,
		Path:     "/",
		HTTPOnly: true,
		Secure:   isProduction,
		SameSite: "None",
		MaxAge:   7 * 24 * 60 * 60,
	}
	c.Cookie(refreshCookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "login successfully",
		"data":    loginResponse,
	})
}

func (h *userHandler) GetProfile(c fiber.Ctx) error {
	data := c.Locals("user")

	claims, ok := data.(*middleware.JWTClaims)
	if !ok {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed get data user",
		})
	}

	userEmail := claims.Email

	user, err := h.authService.GetProfile(userEmail)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successfully get data user",
		"data":    user,
	})
}

func (h *userHandler) RefreshToken(c fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Refresh token required",
		})
	}

	claims, err := h.authService.ValidateRefreshToken(refreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired refresh token",
		})
	}

	userID, ok := claims["userId"]
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token claims",
		})
	}

	userIDStr, ok := userID.(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid userId format",
		})
	}

	userUUID, err := uuid.Parse(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid user ID format",
		})
	}

	user, err := h.employeeService.GetByID(userUUID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	employeeModel := &models.Employee{
		ID:    userUUID,
		Email: user.Email,
	}

	newAccessToken, err := h.authService.GenerateToken(employeeModel)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate new access token",
		})
	}

	newRefreshToken, err := h.authService.GenerateRefreshToken(employeeModel)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate new refresh token",
		})
	}

	isProduction := os.Getenv("ENV") == "production"

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    newAccessToken,
		Path:     "/",
		HTTPOnly: true,
		Secure:   isProduction,
		SameSite: "None",
		MaxAge:   15 * 60,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    newRefreshToken,
		Path:     "/",
		HTTPOnly: true,
		Secure:   isProduction,
		SameSite: "None",
		MaxAge:   7 * 24 * 60 * 60,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "token refreshed successfully",
	})
}

func (h *userHandler) Logout(c fiber.Ctx) error {
	data := c.Locals("user")
	claims, ok := data.(*middleware.JWTClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	userUUID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid user id",
		})
	}

	if err := h.authService.Logout(userUUID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to logout",
		})
	}

	isProduction := os.Getenv("ENV") == "production"

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		HTTPOnly: true,
		Secure:   isProduction,
		SameSite: "Lax",
		MaxAge:   -1,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		HTTPOnly: true,
		Secure:   isProduction,
		SameSite: "Lax",
		MaxAge:   -1,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successfully logout",
	})
}
