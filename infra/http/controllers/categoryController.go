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

type CategoryController struct {
	categoryCommandHandler command.CategoryAction
	categoryQueryHandler   query.CategoryQuery
}

func NewCategoryController(categoryAction *command.CategoryAction, categoryQuery *query.CategoryQuery) *CategoryController {
	return &CategoryController{
		categoryCommandHandler: *categoryAction,
		categoryQueryHandler:   *categoryQuery,
	}
}

func (ctl *CategoryController) FindAll(c *fiber.Ctx) error {
	categories, err := ctl.categoryQueryHandler.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("failed to get categories")
	}

	return c.Status(fiber.StatusOK).JSON(categories)
}

func (ctl *CategoryController) FindById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid parameter")
	}

	categoryId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error in parse id")
	}

	category, err := ctl.categoryQueryHandler.FindById(categoryId)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("failed to get category")
	}

	if category.Description == "" {
		return c.Status(fiber.StatusNotFound).SendString("category not found")
	}

	return c.Status(fiber.StatusOK).JSON(category)
}

func (ctl *CategoryController) Create(c *fiber.Ctx) error {
	validate = validator.New()
	var dtoCategory dto.Category
	err := c.BodyParser(&dtoCategory)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = validate.Struct(dtoCategory)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = ctl.categoryCommandHandler.HandleCreate(dtoCategory)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("failed to create category")
	}

	return c.Status(fiber.StatusOK).SendString("category created successfully")
}

func (ctl *CategoryController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid parameter")
	}

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error in parse id")
	}

	validate = validator.New()
	var categoryDto dto.Category
	err = c.BodyParser(&categoryDto)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("error in parse body request")
	}

	err = validate.Struct(categoryDto)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = ctl.categoryCommandHandler.HandleUpdate(userId, categoryDto)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("error in create category")
	}

	return c.Status(fiber.StatusOK).SendString("category updated successfully")
}

func (ctl *CategoryController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid parameter")
	}

	categoryId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error in parse id")
	}

	err = ctl.categoryCommandHandler.HandleDelete(categoryId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("failed to delete category")
	}

	return c.Status(fiber.StatusOK).SendString("category deleted successfully")
}
