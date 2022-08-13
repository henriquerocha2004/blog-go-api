package controllers

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/blog-go-api/domain/actions/command"
	"github.com/henriquerocha2004/blog-go-api/domain/actions/query"
	"github.com/henriquerocha2004/blog-go-api/infra/http/dto"
)

var validate *validator.Validate

type UserController struct {
	userAction *command.UserAction
	userQuery  *query.UserQuery
}

func NewUserController(userAction *command.UserAction, userQuery *query.UserQuery) *UserController {
	return &UserController{
		userAction: userAction,
		userQuery:  userQuery,
	}
}

func (ctl *UserController) Create(c *fiber.Ctx) error {
	validate = validator.New()
	var userDto dto.UserRequest
	err := c.BodyParser(&userDto)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = validate.Struct(userDto)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err = ctl.userAction.HandleCreate(userDto)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(201).SendString("user created with success!")
}

func (ctl *UserController) FindById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid parameter")
	}

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error in parse id")
	}

	user, err := ctl.userQuery.FindById(userId)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error in search user")
	}

	if user.Id == 0 {
		return c.Status(fiber.StatusNotFound).SendString("user not found")
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (ctl *UserController) FindAll(c *fiber.Ctx) error {
	users, err := ctl.userQuery.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(200).JSON(users)
}

func (ctl *UserController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid parameter")
	}

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error in parse id")
	}

	validate = validator.New()
	var userDto dto.UserRequest
	err = c.BodyParser(&userDto)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = validate.Struct(userDto)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err = ctl.userAction.HandleUpdate(userId, userDto)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error at update user")
	}

	return c.Status(fiber.StatusOK).SendString("user updated successfully")
}

func (ctl *UserController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid parameter")
	}

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error in parse id")
	}

	err = ctl.userAction.HandleDelete(userId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error in delete user")
	}

	return c.Status(fiber.StatusOK).SendString("user deleted successfully")
}
