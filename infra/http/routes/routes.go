package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/blog-go-api/infra/auth"
	"github.com/henriquerocha2004/blog-go-api/infra/container"
)

func Register(app *fiber.App) {
	var di = container.ContainerDependency{}

	api := app.Group("/api")
	api.Post("/auth", di.GetAuthenticateController().Authenticate)
	api.Post("/user", di.GetUserController().Create)

	admin := api.Group("/admin", auth.CheckAuth)
	admin.Get("/users", di.GetUserController().FindAll)
	admin.Get("/user/:id", di.GetUserController().FindById)
	admin.Put("/user/:id", di.GetUserController().Update)
	admin.Delete("/user/:id", di.GetUserController().Delete)
	admin.Post("/category", di.GetCategoryController().Create)
	admin.Put("/category/:id", di.GetCategoryController().Update)
	admin.Delete("/category/:id", di.GetCategoryController().Delete)
	admin.Get("/categories", di.GetCategoryController().FindAll)
	admin.Get("/category/:id", di.GetCategoryController().FindById)
	admin.Post("/post", di.GetPostController().Create)
	admin.Put("/post/:id", di.GetPostController().Update)
	admin.Delete("/post/:id", di.GetPostController().Delete)
	admin.Get("/post/:id", di.GetPostController().FindById)
	admin.Get("/posts/:userId", di.GetPostController().FindByUser)
}
