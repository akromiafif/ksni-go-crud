package route

import (
	"github.com/gofiber/fiber/v2"
	"ksni.com/crud/handler"
)

func RouteInit(r *fiber.App) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"hello": "Oke",
		})
	})
	r.Get("/user", handler.UserHandler)

	r.Listen(":8080")
}