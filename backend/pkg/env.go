package pkg

import (
	"os"
)

type Config struct {
	JWTSecret string
}

var AppConfig Config

func LoadConfig() {
	AppConfig.JWTSecret = os.Getenv("JWT_SECRET")
	if AppConfig.JWTSecret == "" {
		AppConfig.JWTSecret = "your-default-secret-key-change-in-production"
	}
}
