package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/hiteshchouhan22/PlanIt/config"
	"github.com/hiteshchouhan22/PlanIt/db"
	"github.com/hiteshchouhan22/PlanIt/handlers"
	"github.com/hiteshchouhan22/PlanIt/middlewares"
	"github.com/hiteshchouhan22/PlanIt/repositories"
	"github.com/hiteshchouhan22/PlanIt/services"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "Ticket-Booking",
		ServerHeader: "Fiber",
	})

	// Repositories
	eventRepository := repositories.NewEventRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)
	authRepository := repositories.NewAuthRepository(db)

	// Service
	authService := services.NewAuthService(authRepository)

	// Routing
	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	privateRoutes := server.Use(middlewares.AuthProtected(db))

	handlers.NewEventHandler(privateRoutes.Group("/event"), eventRepository)
	handlers.NewTicketHandler(privateRoutes.Group("/ticket"), ticketRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
