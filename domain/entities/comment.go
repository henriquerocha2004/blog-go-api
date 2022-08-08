package domain

type Comment struct {
	Id      int64  `json:"id,omitempty"`
	Content string `json:"content"`
	UserId  int64  `json:"user_id"`
	PostId  int64  `json:"post_id"`
}

type CommentCommand interface {
	Create(comment Comment) error
	Delete(commentId int64) error
	Update(commentId int64, comment Comment) error
}

type CommentQuery interface {
	FindAllByPost(postId int64) ([]Comment, error)
	FindById(commentId int64) (Comment, error)
}
