package routes

import (
	"backend/internal/handlers"
	"backend/internal/repositories"
	"backend/internal/services"
	"backend/pkg"

	"github.com/gofiber/fiber/v3"
)

func EmployeeRoute(router fiber.Router) {
	db := pkg.GetDB()
	if db == nil {
		panic("Database not found")
	}

	pkg.LoadConfig()

	employeeRepository := repositories.NewEmployeeRepository(db)
	userRepository := repositories.NewUserRepository(db)
	employeeService := services.NewEmployeeService(employeeRepository, userRepository, db)
	employeeHandler := handlers.NewEmployeeHandler(employeeService)

	employeeRoute := router.Group("/employee")
	{
		// employeeRoute.Post("/add", middleware.AuthMiddleware(pkg.AppConfig.JWTSecret), employeeHandler.AddNewEmployee)
		employeeRoute.Post("/add", employeeHandler.AddNewEmployee)
	}
}
