package routes

import (
	"backend/internal/handlers"
	"backend/internal/middleware"
	"backend/internal/repositories"
	"backend/internal/services"
	"backend/pkg"

	"github.com/gofiber/fiber/v3"
)

func AuthRouter(router fiber.Router) {
	db := pkg.GetDB()
	if db == nil {
		panic("Database not found")
	}

	pkg.LoadConfig()

	employeeRepository := repositories.NewEmployeeRepository(db)
	userRepository := repositories.NewUserRepository(db)
	authService := services.NewAuthService(employeeRepository, userRepository, pkg.AppConfig.JWTSecret)
	emoployeeService := services.NewEmployeeService(employeeRepository, userRepository, db)
	authHandler := handlers.NewAuthHandler(authService, emoployeeService)

	authGroup := router.Group("/auth")
	{
		authGroup.Post("/login", authHandler.Login)
		authGroup.Get("/profile", middleware.AuthMiddleware(pkg.AppConfig.JWTSecret), authHandler.GetProfile)
		authGroup.Post("/refresh", authHandler.RefreshToken)
		authGroup.Post("/logout", middleware.AuthMiddleware(pkg.AppConfig.JWTSecret), authHandler.Logout)
	}
}
