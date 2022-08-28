package query

import domain "github.com/henriquerocha2004/blog-go-api/domain/entities"

type PostQuery struct {
	query domain.PostQuery
}

func NewPostQuery(postQuery domain.PostQuery) *PostQuery {
	return &PostQuery{
		query: postQuery,
	}
}

func (c *PostQuery) SearchById(postId int64) (domain.Post, error) {
	post, err := c.query.SearchById(postId)
	return *post, err
}

func (c *PostQuery) SearchByUser(userId int64) ([]domain.Post, error) {
	posts, err := c.query.SearchByUser(userId)
	return *posts, err
}
