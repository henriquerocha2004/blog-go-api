package mysql

import (
	"database/sql"
	domain "github.com/henriquerocha2004/blog-go-api/domain/entities"
)

type CommentCommand struct {
	connection *sql.DB
}

type CommentQuery struct {
	connection *sql.DB
}

func NewCommentCommand(connection *sql.DB) *CommentCommand {
	return &CommentCommand{
		connection: connection,
	}
}

func NewCommentQuery(connection *sql.DB) *CommentQuery {
	return &CommentQuery{
		connection: connection,
	}
}

func (c *CommentCommand) Create(comment domain.Comment) error {
	stmt, err := c.connection.Prepare("INSERT INTO comments (content, user_id, post_id) VALUES (?,?,?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(comment.Content, comment.UserId, comment.PostId)
	if err != nil {
		return err
	}
	return nil
}

func (c *CommentCommand) Delete(commentId int64) error {
	stmt, err := c.connection.Prepare("DELETE FROM comments WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(commentId)
	if err != nil {
		return err
	}

	return nil
}

func (c *CommentCommand) Update(commentId int64, comment domain.Comment) error {
	stmt, err := c.connection.Prepare("UPDATE comments SET content = ?, user_id = ?, post_id = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(comment.Content, comment.UserId, comment.PostId, commentId)
	if err != nil {
		return err
	}

	return nil
}

func (q *CommentQuery) FindAllByPost(postId int64) ([]domain.Comment, error) {
	stmt, err := q.connection.Query("SELECT id, content, user_id, post_id FROM comments WHERE post_id = ?", postId)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var comments []domain.Comment

	for stmt.Next() {
		var comment domain.Comment
		stmt.Scan(&comment.Id, &comment.Content, &comment.PostId, &comment.UserId)
		comments = append(comments, comment)
	}

	return comments, nil
}

func (q *CommentQuery) FindById(commentId int64) (*domain.Comment, error) {
	stmt, err := q.connection.Query("SELECT id, content, user_id, post_id FROM comments WHERE id = ?", commentId)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var comment domain.Comment

	for stmt.Next() {
		stmt.Scan(&comment.Id, &comment.Content, &comment.PostId, &comment.UserId)
	}

	return &comment, nil
}
