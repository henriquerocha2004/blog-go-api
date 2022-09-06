package dto

type Comment struct {
	ID      int64  `json:"id,omitempty"`
	Content string `json:"content" validate:"required"`
	UserId  int64  `json:"user_id" validate:"required"`
	PostId  int64  `json:"post_id" validate:"required"`
}
