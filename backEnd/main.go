package main

import (
	"sahih-be/internal/config"
	"sahih-be/internal/database"
	"sahih-be/internal/delivery/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//init config
	viperConfig := config.NewViper()
	envConfig := config.NewEnv(viperConfig)

	//initialize database
	database.DatabaseInit(envConfig)
	//run migrate

	//init server
	app := fiber.New()
	route.RouteInit(app)
	app.Listen(":3000")
}
