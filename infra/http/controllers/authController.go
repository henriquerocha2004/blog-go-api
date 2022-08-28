package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/blog-go-api/infra/auth"
	"github.com/henriquerocha2004/blog-go-api/infra/http/dto"
)

type AuthController struct {
	login *auth.Login
}

func NewAuthController(login *auth.Login) *AuthController {
	return &AuthController{
		login: login,
	}
}

func (ctl *AuthController) Authenticate(ctx *fiber.Ctx) error {
	validate = validator.New()
	var userCredentials dto.UserCredentials
	err := ctx.BodyParser(&userCredentials)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = validate.Struct(userCredentials)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	token, err := ctl.login.Authenticate(userCredentials)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return ctx.Status(fiber.StatusAccepted).JSON(token)
}
