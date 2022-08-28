package mysql

import (
	"database/sql"
	domain "github.com/henriquerocha2004/blog-go-api/domain/entities"
)

type CommentCommand struct {
	connection *sql.DB
}

func NewCommentCommand(connection *sql.DB) *CommentCommand {
	return &CommentCommand{
		connection: connection,
	}
}

func (c *CommentCommand) Create(comment domain.Comment) error {

}
