package dto

type Post struct {
	Id          int64   `json:"id,omitempty"`
	Title       string  `json:"title" validate:"required"`
	Content     string  `json:"content" validate:"required"`
	CategoryIds []int64 `json:"categories_ids" validate:"required"`
	UserId      int64   `json:"user_id" validate:"required"`
}
