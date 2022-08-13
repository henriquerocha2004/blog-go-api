package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/blog-go-api/infra/container"
)

func Register(app *fiber.App) {
	var di container.ContainerDependency = container.ContainerDependency{}
	app.Get("/users", di.GetUserController().FindAll)
	app.Get("/user/:id", di.GetUserController().FindById)
	app.Post("/user", di.GetUserController().Create)
	app.Put("/user/:id", di.GetUserController().Update)
	app.Delete("/user/:id", di.GetUserController().Delete)
}
