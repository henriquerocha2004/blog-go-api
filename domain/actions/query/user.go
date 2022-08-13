package query

import domain "github.com/henriquerocha2004/blog-go-api/domain/entities"

type UserQuery struct {
	query domain.UserQuery
}

func NewUserQuery(userQuery domain.UserQuery) *UserQuery {
	return &UserQuery{
		query: userQuery,
	}
}

func (uq *UserQuery) FindById(userId int64) (domain.User, error) {
	user, err := uq.query.FindById(userId)
	return user, err
}

func (uq *UserQuery) FindAll() ([]domain.User, error) {
	users, err := uq.query.FindAll()
	return *users, err
}
