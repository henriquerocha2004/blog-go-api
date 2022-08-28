package domain

type Post struct {
	Id            int64   `json:"id,omitempty"`
	Title         string  `json:"title"`
	Content       string  `json:"content"`
	CreatedAt     string  `json:"created_at"`
	UserId        int64   `json:"user_id"`
	CategoriesIds []int64 `json:"categories_ids"`
}

type PostCommand interface {
	Create(post Post) error
	Update(postId int64, post Post) error
	Delete(postId int64) error
}

type PostQuery interface {
	SearchById(postId int64) (*Post, error)
	SearchByUser(userId int64) (*[]Post, error)
}
