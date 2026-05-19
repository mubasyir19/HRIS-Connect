package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

type Role string

const (
	RoleAdmin    Role = "admin"
	RoleManager  Role = "manager"
	RoleEmployee Role = "employee"
)

// AllRoles semua role yang tersedia
var AllRoles = []Role{RoleAdmin, RoleManager, RoleEmployee}

// RoleMiddleware membuat middleware untuk mengecek role user
// allowedRoles: daftar role yang diizinkan mengakses endpoint
func RoleMiddleware(allowedRoles ...string) fiber.Handler {
	return func(c fiber.Ctx) error {
		// Ambil role dari context (sudah di-set oleh auth middleware)
		userRole, ok := c.Locals("role").(string)
		if !ok || userRole == "" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error":   "Forbidden",
				"message": "User role not found in context",
			})
		}

		// Cek apakah user memiliki role yang diizinkan
		allowed := false
		for _, role := range allowedRoles {
			if userRole == role {
				allowed = true
				break
			}
		}

		if !allowed {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error":   "Forbidden",
				"message": "You don't have permission to access this resource",
			})
		}

		return c.Next()
	}
}

// RoleMiddlewareWithHierarchy membuat middleware dengan hierarki role
// Admin > Manager > Employee (Admin bisa akses semua, Manager bisa akses Manager & Employee, dll)
func RoleMiddlewareWithHierarchy(minRole string) fiber.Handler {
	roleLevel := map[string]int{
		"admin":    3,
		"manager":  2,
		"employee": 1,
	}

	minLevel, exists := roleLevel[minRole]
	if !exists {
		minLevel = 0
	}

	return func(c fiber.Ctx) error {
		userRole, ok := c.Locals("role").(string)
		if !ok || userRole == "" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error":   "Forbidden",
				"message": "User role not found",
			})
		}

		userLevel, exists := roleLevel[userRole]
		if !exists || userLevel < minLevel {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error":   "Forbidden",
				"message": "Insufficient permissions",
			})
		}

		return c.Next()
	}
}

// ResourceOwnerMiddleware middleware untuk mengecek apakah user adalah pemilik resource
// getResourceEmployeeID: function untuk mendapatkan employee ID dari resource
// getResourceUserID: function untuk mendapatkan user ID dari resource (opsional)
func ResourceOwnerMiddleware(getResourceEmployeeID func(c fiber.Ctx) (string, error)) fiber.Handler {
	return func(c fiber.Ctx) error {
		// Ambil user ID dan role dari context
		userID, err := GetUserID(c)
		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error":   "Forbidden",
				"message": "User ID not found",
			})
		}

		userRole := GetUserRole(c)

		// Admin bisa akses semua resource
		if userRole == string(RoleAdmin) {
			return c.Next()
		}

		// Untuk role selain admin, cek apakah dia pemilik resource
		resourceEmployeeID, err := getResourceEmployeeID(c)
		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error":   "Forbidden",
				"message": "Resource not found or invalid",
			})
		}

		// Parse resource employee ID
		resourceID, err := parseUUID(resourceEmployeeID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "Bad Request",
				"message": "Invalid resource ID format",
			})
		}

		// Konversi resourceID ke string untuk perbandingan
		resourceIDStr := resourceID

		// Manager bisa akses resource bawahannya (perlu logic tambahan di service)
		if userRole == string(RoleManager) {
			// Ini akan diimplementasikan di service layer
			// Manager access check dilakukan di service/controller
			c.Locals("resource_owner_check", resourceIDStr)
			c.Locals("user_id", userID)
			return c.Next()
		}

		// Employee hanya bisa akses resource milik sendiri
		// Bandingkan user ID dengan resource employee ID
		employeeIDFromUser, err := getEmployeeIDFromUserID(c, userID)
		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error":   "Forbidden",
				"message": "Cannot verify resource ownership",
			})
		}

		if employeeIDFromUser != resourceIDStr {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error":   "Forbidden",
				"message": "You can only access your own resources",
			})
		}

		return c.Next()
	}
}

// HasRole helper function untuk mengecek apakah user memiliki role tertentu di dalam handler
func HasRole(c fiber.Ctx, targetRole string) bool {
	userRole := GetUserRole(c)
	return userRole == targetRole
}

// HasAnyRole helper function untuk mengecek apakah user memiliki minimal satu role yang diizinkan
func HasAnyRole(c fiber.Ctx, roles ...string) bool {
	userRole := GetUserRole(c)
	for _, role := range roles {
		if userRole == role {
			return true
		}
	}
	return false
}

// IsAdmin helper function
func IsAdmin(c fiber.Ctx) bool {
	return HasRole(c, string(RoleAdmin))
}

// IsManager helper function
func IsManager(c fiber.Ctx) bool {
	return HasRole(c, string(RoleManager))
}

// IsEmployee helper function
func IsEmployee(c fiber.Ctx) bool {
	return HasRole(c, string(RoleEmployee))
}

// GetMinRoleFromEndpoint mendapatkan role minimal untuk sebuah endpoint
// Bisa digunakan untuk dynamic permission checking
func GetMinRoleFromEndpoint(path, method string) string {
	// Mapping endpoint ke minimal role
	// Bisa disimpan di database untuk dynamic permission
	permissions := map[string]map[string]string{
		"/api/v1/employees": {
			"GET":    "manager",
			"POST":   "admin",
			"PUT":    "admin",
			"DELETE": "admin",
		},
		"/api/v1/leave-requests": {
			"GET":  "employee",
			"POST": "employee",
		},
		"/api/v1/leave-requests/pending": {
			"GET": "manager",
		},
		"/api/v1/dashboard": {
			"GET": "manager",
		},
		"/api/v1/settings": {
			"GET":  "admin",
			"PUT":  "admin",
			"POST": "admin",
		},
	}

	if endpointPerms, ok := permissions[path]; ok {
		if role, ok := endpointPerms[method]; ok {
			return role
		}
	}

	// Default: hanya admin yang bisa akses
	return "admin"
}

// Helper function untuk parse UUID
func parseUUID(id string) (string, error) {
	// Validasi UUID format (basic check)
	if len(id) == 0 {
		return "", fiber.NewError(fiber.StatusBadRequest, "Invalid UUID format")
	}
	return id, nil
}

// Helper function untuk mendapatkan employee ID dari user ID
func getEmployeeIDFromUserID(c fiber.Ctx, userID interface{}) (string, error) {
	// Ini perlu dipanggil ke service layer
	// Placeholder implementation - convert userID to string
	if userID == nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, "User ID is nil")
	}

	// Try to convert to string
	switch v := userID.(type) {
	case string:
		return v, nil
	default:
		return fmt.Sprintf("%v", userID), nil
	}
}
