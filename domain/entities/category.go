package domain

type Category struct {
	Id          int64  `json:"id,omitempty"`
	Description string `json:"description"`
}

type CategoryCommand interface {
	Create(category Category) error
	Delete(categoryId int64) error
	Update(categoryId int64, category Category) error
}

type CategoryQuery interface {
	FindAll() ([]Category, error)
	FindById(categoryId int64) (Category, error)
	FindByPostId(postId int64) (Category, error)
}
