package main

import (
	"acne-scan-api/configs"
	"acne-scan-api/internal/app"
	"acne-scan-api/internal/infrastructure/database"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	fiber := fiber.New()

	config, err := configs.LoadConfig()
	if err != nil {
		logrus.Fatal("Error loading config:", err.Error())
	}

	// Initialize database connection
	db, err := database.NewMySQLConnection(&config.MySQL)
	if err != nil {
		logrus.Fatal("Error connecting to MySQL:", err.Error())
	}

	validate := validator.New()

	app.InitApp(db,validate,fiber)

	fiber.Use(cors.New())
	fiber.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	fiber.Listen(":8080")
}