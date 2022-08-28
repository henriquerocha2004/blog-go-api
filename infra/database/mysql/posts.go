package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	domain "github.com/henriquerocha2004/blog-go-api/domain/entities"
)

type PostCommand struct {
	connection *sql.DB
}

type PostQuery struct {
	connection *sql.DB
}

func NewPostCommand(connection *sql.DB) *PostCommand {
	return &PostCommand{
		connection: connection,
	}
}

func NewPostQuery(connection *sql.DB) *PostQuery {
	return &PostQuery{
		connection: connection,
	}
}

func (q *PostQuery) SearchById(postId int64) (*domain.Post, error) {
	stmt, err := q.connection.Query("SELECT id, title, content, user_id, created_at FROM posts WHERE id = ?", postId)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var post domain.Post

	for stmt.Next() {
		err = stmt.Scan(&post.Id, &post.Title, &post.Content, &post.UserId, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &post, nil
}

func (q *PostQuery) SearchByUser(userId int64) (*[]domain.Post, error) {
	stmt, err := q.connection.Query("SELECT id, title, content, user_id, created_at FROM posts WHERE user_id = ?", userId)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var posts []domain.Post

	for stmt.Next() {
		var post domain.Post
		err = stmt.Scan(&post.Id, &post.Title, &post.Content, &post.UserId, &post.CreatedAt)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return &posts, nil

}

func (c *PostCommand) Create(post domain.Post) error {

	ctx := context.Background()
	stmt, err := c.connection.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	row, err := stmt.PrepareContext(ctx, "INSERT INTO posts (title, content, created_at, user_id) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}

	defer row.Close()

	result, err := row.ExecContext(
		ctx,
		post.Title,
		post.Content,
		time.Now().Format("2006-01-02 15:04"),
		post.UserId,
	)

	if err != nil {
		_ = stmt.Rollback()
		return err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		_ = stmt.Rollback()
		return err
	}

	row, err = stmt.PrepareContext(ctx, "INSERT INTO category_post (category_id, post_id) VALUES (?,?)")
	if err != nil {
		_ = stmt.Rollback()
		return err
	}

	for _, categoryId := range post.CategoriesIds {
		_, err = row.ExecContext(ctx, categoryId, lastInsertId)
		if err != nil {
			_ = stmt.Rollback()
			return err
		}
	}

	err = stmt.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (c *PostCommand) Update(postId int64, post domain.Post) error {
	ctx := context.Background()
	stmt, err := c.connection.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	row, err := stmt.PrepareContext(ctx, "UPDATE posts SET title = ?, content = ?, user_id = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer func(row *sql.Stmt) {
		err := row.Close()
		if err != nil {

		}
	}(row)

	_, err = row.ExecContext(
		ctx,
		post.Title,
		post.Content,
		post.UserId,
		postId,
	)

	if err != nil {
		return err
	}

	// SYNC Categories
	row, err = stmt.PrepareContext(ctx, "DELETE FROM category_post WHERE post_id = ?")
	if err != nil {
		fmt.Println("error in delete category post")
		stmt.Rollback()
		return err
	}

	_, err = row.ExecContext(ctx, postId)
	if err != nil {
		fmt.Println("error in delete category post")
		_ = stmt.Rollback()
		return err
	}

	row, err = stmt.PrepareContext(ctx, "INSERT INTO category_post (category_id, post_id) VALUES (?,?)")
	if err != nil {
		_ = stmt.Rollback()
		return err
	}

	for _, category := range post.CategoriesIds {
		_, err = row.ExecContext(ctx, category, postId)
		if err != nil {
			_ = stmt.Rollback()
			return err
		}
	}

	err = stmt.Commit()

	if err != nil {
		return err
	}

	return nil
}

func (c *PostCommand) Delete(postId int64) error {
	ctx := context.Background()
	stmt, err := c.connection.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	row, err := stmt.PrepareContext(ctx, "DELETE FROM category_post WHERE post_id = ?")
	if err != nil {
		_ = stmt.Rollback()
		return err
	}

	_, err = row.ExecContext(ctx, postId)
	if err != nil {
		_ = stmt.Rollback()
		return err
	}

	row, err = stmt.PrepareContext(ctx, "DELETE FROM posts WHERE id = ?")
	if err != nil {
		_ = stmt.Rollback()
		return err
	}

	_, err = row.ExecContext(ctx, postId)
	if err != nil {
		_ = stmt.Rollback()
		return err
	}

	err = stmt.Commit()
	if err != nil {
		return err
	}

	return nil
}
