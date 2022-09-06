package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/blog-go-api/domain/actions/command"
	"github.com/henriquerocha2004/blog-go-api/domain/actions/query"
	"github.com/henriquerocha2004/blog-go-api/infra/http/dto"
	"strconv"
)

type CommentController struct {
	actionComment *command.CommentAction
	queryComment  *query.CommentQuery
}

func NewCommentController(actionComment *command.CommentAction, queryCommand *query.CommentQuery) *CommentController {
	return &CommentController{
		actionComment: actionComment,
		queryComment:  queryCommand,
	}
}

func (ctl *CommentController) Create(ctx *fiber.Ctx) error {
	validate = validator.New()
	var dtoComment dto.Comment
	err := ctx.BodyParser(&dtoComment)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = validate.Struct(dtoComment)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = ctl.actionComment.HandleCreate(dtoComment)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("error in create comment")
	}

	return ctx.Status(fiber.StatusOK).SendString("comment created with successfully")
}

func (ctl *CommentController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid parameter")
	}

	commentId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error in parse id")
	}

	validate = validator.New()
	var commentDto dto.Comment
	err = c.BodyParser(&commentDto)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("error in parse body request")
	}

	err = ctl.actionComment.HandleUpdate(commentId, commentDto)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("error in update comment")
	}

	return c.Status(fiber.StatusAccepted).SendString("comment updated with successfully")
}

func (ctl *CommentController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid parameter")
	}

	commentId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error in parse id")
	}

	err = ctl.actionComment.HandleDelete(commentId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("error in delete comment")
	}

	return c.Status(fiber.StatusOK).SendString("comment deleted with successfully")
}

func (ctl *CommentController) SearchByPost(c *fiber.Ctx) error {
	id := c.Params("postId")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid parameter")
	}

	postId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error in parse id")
	}

	comments, err := ctl.queryComment.SearchByPost(postId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("failed to search comments")
	}

	return c.Status(fiber.StatusOK).JSON(comments)
}

func (ctl *CommentController) SearchById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid parameter")
	}

	commentId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error in parse id")
	}

	comment, err := ctl.queryComment.SearchById(commentId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("failed to search comment")
	}

	return c.Status(fiber.StatusOK).JSON(comment)
}
