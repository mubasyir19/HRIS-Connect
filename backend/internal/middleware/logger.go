package middleware

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
)

// LogLevel level logging
type LogLevel string

const (
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
	LogLevelDebug LogLevel = "debug"
)

// LoggerConfig konfigurasi untuk logger middleware
type LoggerConfig struct {
	// SkipPaths daftar path yang tidak perlu di-log
	SkipPaths []string

	// SkipBody untuk body request/response (untuk keamanan)
	SkipBody bool

	// LogHeaders daftar header yang akan di-log
	LogHeaders []string

	// LogLevel minimal level yang akan di-log
	LogLevel LogLevel

	// Output writer untuk log
	Output io.Writer

	// IncludeResponseBody apakah menyertakan response body
	IncludeResponseBody bool

	// IncludeRequestBody apakah menyertakan request body
	IncludeRequestBody bool

	// MaxBodySize maksimal ukuran body yang di-log
	MaxBodySize int
}

// DefaultLoggerConfig mengembalikan konfigurasi default
func DefaultLoggerConfig() LoggerConfig {
	return LoggerConfig{
		SkipPaths:           []string{"/health", "/metrics", "/api/v1/auth/login"},
		SkipBody:            false,
		LogHeaders:          []string{"User-Agent", "Content-Type", "Authorization"},
		LogLevel:            LogLevelInfo,
		Output:              os.Stdout,
		IncludeResponseBody: false,
		IncludeRequestBody:  true,
		MaxBodySize:         1024, // 1KB
	}
}

// LoggerMiddleware membuat middleware untuk logging request/response
func LoggerMiddleware() fiber.Handler {
	return LoggerMiddlewareWithConfig(DefaultLoggerConfig())
}

// LoggerMiddlewareWithConfig membuat logger middleware dengan konfigurasi custom
func LoggerMiddlewareWithConfig(config LoggerConfig) fiber.Handler {
	return func(c fiber.Ctx) error {
		// Skip logging untuk path tertentu
		for _, skipPath := range config.SkipPaths {
			if strings.HasPrefix(c.Path(), skipPath) {
				return c.Next()
			}
		}

		// Mulai timer
		start := time.Now()

		// Baca request body
		var requestBody string
		if config.IncludeRequestBody && !config.SkipBody {
			requestBody = readBody(c, config.MaxBodySize)
			// Restore body untuk digunakan handler
			c.Request().SetBody([]byte(requestBody))
		}

		// Note: Response body capture tidak didukung sepenuhnya di Fiber v3
		// Response sudah ditulis saat ini, jadi tidak bisa di-capture setelah Next()

		// Proses request ke handler berikutnya
		err := c.Next()

		// Hitung durasi
		duration := time.Since(start)

		// Prepare log entry
		logEntry := LogEntry{
			Timestamp:   time.Now(),
			Method:      c.Method(),
			Path:        c.Path(),
			Query:       string(c.Request().URI().QueryString()),
			IP:          c.IP(),
			UserAgent:   c.Get("User-Agent"),
			StatusCode:  c.Response().StatusCode(),
			Duration:    duration,
			RequestBody: requestBody,
			Headers:     extractHeaders(c, config.LogHeaders),
			Error:       err,
		}

		// Response body tidak disertakan (Fiber v3 tidak support reading written response)

		// Log berdasarkan status code
		statusCode := c.Response().StatusCode()
		level := getLogLevelFromStatusCode(statusCode)

		// Cek apakah level yang akan di-log memenuhi kriteria
		if shouldLog(level, config.LogLevel) {
			logEntry.Log(level, config.Output)
		}

		return err
	}
}

// LogEntry struktur log entry
type LogEntry struct {
	Timestamp    time.Time         `json:"timestamp"`
	Method       string            `json:"method"`
	Path         string            `json:"path"`
	Query        string            `json:"query,omitempty"`
	IP           string            `json:"ip"`
	UserAgent    string            `json:"user_agent"`
	StatusCode   int               `json:"status_code"`
	Duration     time.Duration     `json:"duration_ms"`
	RequestBody  string            `json:"request_body,omitempty"`
	ResponseBody string            `json:"response_body,omitempty"`
	Headers      map[string]string `json:"headers,omitempty"`
	Error        error             `json:"error,omitempty"`
}

// Log menulis log entry ke output
func (e *LogEntry) Log(level LogLevel, output io.Writer) {
	// Format JSON untuk structured logging
	logData := map[string]interface{}{
		"timestamp":   e.Timestamp.Format(time.RFC3339Nano),
		"level":       level,
		"method":      e.Method,
		"path":        e.Path,
		"ip":          e.IP,
		"status_code": e.StatusCode,
		"duration_ms": e.Duration.Milliseconds(),
	}

	if e.Query != "" {
		logData["query"] = e.Query
	}
	if e.UserAgent != "" {
		logData["user_agent"] = e.UserAgent
	}
	if e.RequestBody != "" && level == LogLevelDebug {
		logData["request_body"] = e.RequestBody
	}
	if e.ResponseBody != "" && level == LogLevelDebug {
		logData["response_body"] = e.ResponseBody
	}
	if e.Error != nil {
		logData["error"] = e.Error.Error()
	}
	if len(e.Headers) > 0 {
		// Hapus sensitive headers
		delete(e.Headers, "Authorization")
		delete(e.Headers, "authorization")
		delete(e.Headers, "Cookie")
		delete(e.Headers, "cookie")
		logData["headers"] = e.Headers
	}

	jsonData, _ := json.Marshal(logData)
	log.Println(string(jsonData))
}

// readBody membaca body request dengan aman
func readBody(c fiber.Ctx, maxSize int) string {
	body := c.Body()
	if len(body) == 0 {
		return ""
	}

	// Batasi ukuran body
	if len(body) > maxSize {
		return string(body[:maxSize]) + "... (truncated)"
	}

	return string(body)
}

// extractHeaders mengekstrak header yang diperlukan
func extractHeaders(c fiber.Ctx, headers []string) map[string]string {
	result := make(map[string]string)
	for _, h := range headers {
		if value := c.Get(h); value != "" {
			result[h] = value
		}
	}
	return result
}

// getLogLevelFromStatusCode menentukan level log berdasarkan status code
func getLogLevelFromStatusCode(statusCode int) LogLevel {
	switch {
	case statusCode >= 500:
		return LogLevelError
	case statusCode >= 400:
		return LogLevelWarn
	default:
		return LogLevelInfo
	}
}

// shouldLog mengecek apakah level saat ini perlu di-log
func shouldLog(currentLevel, minLevel LogLevel) bool {
	levels := map[LogLevel]int{
		LogLevelDebug: 1,
		LogLevelInfo:  2,
		LogLevelWarn:  3,
		LogLevelError: 4,
	}

	return levels[currentLevel] >= levels[minLevel]
}

// GinStyleLogger middleware dengan format seperti Gin framework
func GinStyleLogger() fiber.Handler {
	return func(c fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		duration := time.Since(start)
		statusCode := c.Response().StatusCode()

		// Warna untuk status code
		color := "\033[32m" // Green (200)
		if statusCode >= 400 {
			color = "\033[31m" // Red
		} else if statusCode >= 300 {
			color = "\033[33m" // Yellow
		}

		log.Printf("[%s] %s %s %d %s %v",
			time.Now().Format("2006-01-02 15:04:05"),
			c.Method(),
			c.Path(),
			statusCode,
			color,
			duration,
		)

		return err
	}
}

// AuditLogger middleware khusus untuk logging operasi critical (create, update, delete)
func AuditLogger() fiber.Handler {
	return func(c fiber.Ctx) error {
		// Hanya log untuk method yang mengubah data
		method := c.Method()
		if method != "POST" && method != "PUT" && method != "DELETE" && method != "PATCH" {
			return c.Next()
		}

		// Ambil user info dari context (sudah di-set oleh auth middleware)
		userID, err := GetUserID(c)
		if err != nil {
			// User tidak terautentikasi, skip audit log
			return c.Next()
		}

		userRole := GetUserRole(c)

		// Baca request body
		requestBody := readBody(c, 2048)

		// Proses request
		reqErr := c.Next()

		// Log audit
		auditLog := map[string]interface{}{
			"timestamp":    time.Now().Format(time.RFC3339),
			"user_id":      userID.String(),
			"user_role":    userRole,
			"method":       method,
			"path":         c.Path(),
			"status_code":  c.Response().StatusCode(),
			"request_body": requestBody,
		}

		if reqErr != nil {
			auditLog["error"] = reqErr.Error()
		}

		jsonData, _ := json.Marshal(auditLog)
		log.Printf("[AUDIT] %s", string(jsonData))

		return reqErr
	}
}
