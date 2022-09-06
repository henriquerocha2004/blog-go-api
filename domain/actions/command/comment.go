package command

import (
	domain "github.com/henriquerocha2004/blog-go-api/domain/entities"
	"github.com/henriquerocha2004/blog-go-api/infra/http/dto"
)

type CommentAction struct {
	commentAction domain.CommentCommand
}

func NewCommentAction(commentAction domain.CommentCommand) *CommentAction {
	return &CommentAction{
		commentAction: commentAction,
	}
}

func (c *CommentAction) HandleCreate(dtoComment dto.Comment) error {

	comment := domain.Comment{
		UserId:  dtoComment.UserId,
		PostId:  dtoComment.PostId,
		Content: dtoComment.Content,
	}

	err := c.commentAction.Create(comment)
	if err != nil {
		return err
	}
	return nil
}

func (c *CommentAction) HandleUpdate(commentId int64, dtoComment dto.Comment) error {
	comment := domain.Comment{
		UserId:  dtoComment.UserId,
		PostId:  dtoComment.PostId,
		Content: dtoComment.Content,
	}

	err := c.commentAction.Update(commentId, comment)
	if err != nil {
		return err
	}
	return nil
}

func (c *CommentAction) HandleDelete(commentId int64) error {
	return c.commentAction.Delete(commentId)
}
