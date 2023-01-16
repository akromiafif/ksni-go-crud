package main

import (
	"github.com/gofiber/fiber/v2"
	"ksni.com/crud/database"
	"ksni.com/crud/route"
)

func main() {
	app := fiber.New()

	// INIT DATABASE
	database.DatabaseInit()

	// INIT ROUTE
	route.RouteInit(app);

	app.Listen(":8080")
}