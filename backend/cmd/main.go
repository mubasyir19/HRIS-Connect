package main

import (
	"backend/internal/routes"
	"backend/pkg"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	if err := pkg.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// db := pkg.GetDB()
	// if db == nil {
	// 	panic("Database not found")
	// }
	// if err := pkg.MigrateDB(db); err != nil {
	// 	log.Fatal("Failed to migrate database:", err)
	// }
	// log.Println("✅ Database migration completed")

	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Access-Token"},
	}))

	apiV1 := app.Group("/api/v1")

	routes.AuthRouter(apiV1)
	routes.EmployeeRoute(apiV1)
	routes.DepartmentRouter(apiV1)

	apiV1.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World")
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on http://localhost:%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
