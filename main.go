package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"go-api-template/handlers"
	"go-api-template/repositories"
	s "go-api-template/services"
	"log"
	"os"
	"strings"
)

func main() {
	verifyEnvs()
	repos := repositories.InitRepositories()
	services := s.InitServices(repos)
	handler := handlers.NewHandler(repos, services)

	app := fiber.New(fiber.Config{
		ErrorHandler: handler.Error,
	})
	app.Use(recover.New())
	app.Use(cors.New(corsConfig()))
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	initRoutes(app, handler)

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func corsConfig() cors.Config {
	headers := []string{
		"Origin",
		"Content-Type",
		"Accept",
		"X-Auth-Token",
		"X-version",
		"Access-Control-Allow-Origin",
	}
	return cors.Config{
		AllowOrigins: strings.Split(os.Getenv("Origins"), ","),
		AllowHeaders: headers,
		AllowMethods: []string{"GET", "POST", "HEAD", "PUT", "DELETE"},
		MaxAge:       86400,
	}
}
