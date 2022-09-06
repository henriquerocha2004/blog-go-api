package query

import domain "github.com/henriquerocha2004/blog-go-api/domain/entities"

type CommentQuery struct {
	commentQuery domain.CommentQuery
}

func NewCommentQuery(commentQuery domain.CommentQuery) *CommentQuery {
	return &CommentQuery{
		commentQuery: commentQuery,
	}
}

func (q *CommentQuery) SearchByPost(postId int64) ([]domain.Comment, error) {
	comments, err := q.commentQuery.FindAllByPost(postId)
	return comments, err
}

func (q *CommentQuery) SearchById(postId int64) (*domain.Comment, error) {
	comment, err := q.commentQuery.FindById(postId)
	return comment, err
}
