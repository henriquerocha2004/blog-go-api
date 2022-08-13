package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/blog-go-api/infra/http/routes"
)

func main() {
	app := fiber.New()
	routes.Register(app)
	app.Listen(":8000")
}
