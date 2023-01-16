package route

import (
	"github.com/gofiber/fiber/v2"
	"ksni.com/crud/handler"
)

func RouteInit(r *fiber.App) {
	r.Get("/user", handler.UserHandlerGetAll)
	r.Post("/user", handler.UserHandlerCreate)
	r.Get("/user/:id", handler.UserHandlerGetById)

	r.Listen("127.0.0.1:8080")
}