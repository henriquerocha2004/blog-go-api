package dto

type Category struct {
	Id          int    `json:"id,omitempty"`
	Description string `json:"description" validate:"required"`
}
