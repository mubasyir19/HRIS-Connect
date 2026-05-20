package routes

import (
	"backend/internal/handlers"
	"backend/internal/middleware"
	"backend/internal/repositories"
	"backend/internal/services"
	"backend/pkg"

	"github.com/gofiber/fiber/v3"
)

func DepartmentRouter(router fiber.Router) {
	db := pkg.GetDB()
	if db == nil {
		panic("Database not found")
	}

	pkg.LoadConfig()

	departmentRepository := repositories.NewDepartmentRepository(db)
	employeeRepository := repositories.NewEmployeeRepository(db)
	departmentService := services.NewDepartmentService(departmentRepository, employeeRepository)
	departmentHandler := handlers.NewDepartmentHandler(departmentService)

	departmentGroup := router.Group("/department")
	{
		departmentGroup.Get("/all", middleware.AuthMiddleware(pkg.AppConfig.JWTSecret), departmentHandler.GetListAllDepartment)
		departmentGroup.Post("/add", middleware.AuthMiddleware(pkg.AppConfig.JWTSecret), departmentHandler.AddNewDepartment)
	}
}
