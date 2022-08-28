package controllers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/blog-go-api/domain/actions/command"
	"github.com/henriquerocha2004/blog-go-api/domain/actions/query"
	"github.com/henriquerocha2004/blog-go-api/infra/http/dto"
	"strconv"
)

type PostController struct {
	actionCommand *command.PostAction
	query         *query.PostQuery
}

func NewPostController(actionCommand *command.PostAction, query *query.PostQuery) *PostController {
	return &PostController{
		actionCommand: actionCommand,
		query:         query,
	}
}

func (ctl *PostController) Create(ctx *fiber.Ctx) error {
	validate = validator.New()
	var dtoPost dto.Post
	err := ctx.BodyParser(&dtoPost)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = validate.Struct(dtoPost)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = ctl.actionCommand.HandleCreate(dtoPost)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("error at create post")
	}

	return ctx.Status(fiber.StatusOK).SendString("post created successfully")
}

func (ctl *PostController) Update(ctx *fiber.Ctx) error {
	validate = validator.New()
	var dtoPost dto.Post
	err := ctx.BodyParser(&dtoPost)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid parameter")
	}

	postId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error in parse id")
	}

	err = validate.Struct(dtoPost)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = ctl.actionCommand.HandleUpdate(postId, dtoPost)

	if err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("error in update post")
	}

	return ctx.Status(fiber.StatusOK).SendString("post updated successfully")
}

func (ctl *PostController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid parameter")
	}

	postId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error in parse id")
	}

	err = ctl.actionCommand.HandleDelete(postId)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("error in delete post")
	}

	return ctx.Status(fiber.StatusOK).SendString("post deleted successfully")
}

func (ctl *PostController) FindById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid parameter")
	}

	postId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error in parse id")
	}

	post, err := ctl.query.SearchById(postId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("failed to get post")
	}

	if post.Id == 0 {
		return ctx.Status(fiber.StatusNotFound).SendString("post not found")
	}

	return ctx.Status(fiber.StatusOK).JSON(post)
}

func (ctl *PostController) FindByUser(ctx *fiber.Ctx) error {
	id := ctx.Params("userId")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid parameter")
	}

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error in parse id")
	}

	posts, err := ctl.query.SearchByUser(userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("failed to retrieve posts by user")
	}

	return ctx.Status(fiber.StatusOK).JSON(posts)
}
