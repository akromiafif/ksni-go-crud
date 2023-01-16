package main

import (
	"github.com/gofiber/fiber/v2"
	"ksni.com/crud/database"
	"ksni.com/crud/database/migration"
	"ksni.com/crud/route"
)

func main() {
	app := fiber.New()

	// INIT DATABASE
	database.DatabaseInit()

	// RUN MIGRATION
	migration.RunMigration()

	// INIT ROUTE
	route.RouteInit(app);

	app.Listen("127.0.0.1:8080")
}