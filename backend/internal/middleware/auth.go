package middleware

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// JWTClaims struct untuk claims JWT token
type JWTClaims struct {
	UserID     string `json:"userId"`
	EmployeeID string `json:"employeeId"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	jwt.RegisteredClaims
}

// AuthConfig konfigurasi untuk auth middleware
type AuthConfig struct {
	JWTSecret        string
	TokenLookup      string // header:Authorization, query:token, cookie:token
	AuthScheme       string // Bearer
	ContextKeyUser   string // key untuk menyimpan user di context
	ContextKeyUserID string // key untuk menyimpan user ID di context
	ContextKeyRole   string // key untuk menyimpan role di context
}

// DefaultAuthConfig mengembalikan konfigurasi default
func DefaultAuthConfig(jwtSecret string) AuthConfig {
	return AuthConfig{
		JWTSecret: jwtSecret,
		// TokenLookup:      "header:Authorization",
		// AuthScheme:       "Bearer",
		TokenLookup:      "cookie:access_token",
		AuthScheme:       "",
		ContextKeyUser:   "user",
		ContextKeyUserID: "userId",
		ContextKeyRole:   "role",
	}
}

// AuthMiddleware membuat middleware untuk autentikasi JWT
func AuthMiddleware(jwtSecret string) fiber.Handler {
	return AuthMiddlewareWithConfig(DefaultAuthConfig(jwtSecret))
}

// AuthMiddlewareWithConfig membuat middleware autentikasi dengan konfigurasi custom
func AuthMiddlewareWithConfig(config AuthConfig) fiber.Handler {
	return func(c fiber.Ctx) error {
		// Extract token dari request
		token, err := extractToken(c, config)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
		}

		// Parse dan validasi token
		claims, err := parseAndValidateToken(token, config.JWTSecret)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "Unauthorized",
				"message": "Invalid or expired token",
			})
		}

		// Simpan data user ke context Fiber Locals
		c.Locals(config.ContextKeyUser, claims)
		c.Locals(config.ContextKeyUserID, claims.UserID)
		c.Locals(config.ContextKeyRole, claims.Role)

		// Lanjut ke handler berikutnya
		return c.Next()
	}
}

// extractToken mengekstrak token dari berbagai sumber (header, query, cookie)
func extractToken(c fiber.Ctx, config AuthConfig) (string, error) {
	// Parse lookup format (contoh: "header:Authorization")
	parts := strings.SplitN(config.TokenLookup, ":", 2)
	if len(parts) != 2 {
		return "", fiber.NewError(fiber.StatusInternalServerError, "Invalid token lookup configuration")
	}

	source := parts[0] // header, query, cookie
	name := parts[1]   // Authorization, token, dll

	var token string

	switch source {
	case "header":
		token = c.Get(name)
		token = strings.TrimPrefix(token, config.AuthScheme+" ")
		token = strings.TrimSpace(token)
	case "query":
		token = c.Query(name)
	case "cookie":
		token = c.Cookies(name)
	default:
		return "", fiber.NewError(fiber.StatusUnauthorized, "Unsupported token source")
	}

	if token == "" {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Missing authentication token")
	}

	return token, nil
}

// parseAndValidateToken memparsing dan memvalidasi JWT token
func parseAndValidateToken(tokenString, jwtSecret string) (*JWTClaims, error) {
	claims := &JWTClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Validasi signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	// Cek apakah token expired
	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Token expired")
	}

	return claims, nil
}

// GetUserID helper function untuk mendapatkan user ID dari context
func GetUserID(c fiber.Ctx) (uuid.UUID, error) {
	userID, ok := c.Locals("userId").(string)
	if !ok || userID == "" {
		return uuid.Nil, fiber.NewError(fiber.StatusUnauthorized, "User ID not found in context")
	}

	id, err := uuid.Parse(userID)
	if err != nil {
		return uuid.Nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid user ID format")
	}

	return id, nil
}

// GetUserRole helper function untuk mendapatkan role dari context
func GetUserRole(c fiber.Ctx) string {
	role, ok := c.Locals("role").(string)
	if !ok {
		return ""
	}
	return role
}

// GetClaims helper function untuk mendapatkan claims dari context
func GetClaims(c fiber.Ctx) *JWTClaims {
	claims, ok := c.Locals("user").(*JWTClaims)
	if !ok {
		return nil
	}
	return claims
}

// OptionalAuthMiddleware middleware opsional (token tidak wajib)
// Jika token ada dan valid, simpan user ke context, jika tidak tetap lanjutkan
func OptionalAuthMiddleware(jwtSecret string) fiber.Handler {
	return func(c fiber.Ctx) error {
		config := DefaultAuthConfig(jwtSecret)

		token, err := extractToken(c, config)
		if err != nil {
			// Token tidak ada atau invalid, tetap lanjutkan tanpa user context
			return c.Next()
		}

		claims, err := parseAndValidateToken(token, jwtSecret)
		if err != nil {
			// Token invalid, tetap lanjutkan tanpa user context
			return c.Next()
		}

		// Token valid, simpan user ke context
		c.Locals(config.ContextKeyUser, claims)
		c.Locals(config.ContextKeyUserID, claims.UserID)
		c.Locals(config.ContextKeyRole, claims.Role)

		return c.Next()
	}
}
