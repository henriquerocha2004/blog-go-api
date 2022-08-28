package command

import (
	domain "github.com/henriquerocha2004/blog-go-api/domain/entities"
	"github.com/henriquerocha2004/blog-go-api/infra/http/dto"
)

type PostAction struct {
	postCommand domain.PostCommand
}

func NewPostAction(postCommand domain.PostCommand) *PostAction {
	return &PostAction{
		postCommand: postCommand,
	}
}

func (c *PostAction) HandleCreate(dtoPost dto.Post) error {

	post := domain.Post{
		Title:         dtoPost.Title,
		Content:       dtoPost.Content,
		UserId:        dtoPost.UserId,
		CategoriesIds: dtoPost.CategoryIds,
	}

	err := c.postCommand.Create(post)

	if err != nil {
		return err
	}

	return nil
}

func (c *PostAction) HandleUpdate(postId int64, dtoPost dto.Post) error {
	post := domain.Post{
		Title:         dtoPost.Title,
		Content:       dtoPost.Content,
		UserId:        dtoPost.UserId,
		CategoriesIds: dtoPost.CategoryIds,
	}

	err := c.postCommand.Update(postId, post)

	if err != nil {
		return err
	}

	return nil
}

func (c *PostAction) HandleDelete(postId int64) error {
	return c.postCommand.Delete(postId)
}
