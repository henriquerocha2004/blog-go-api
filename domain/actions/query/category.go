package query

import (
	domain "github.com/henriquerocha2004/blog-go-api/domain/entities"
)

type CategoryQuery struct {
	categoryQuery domain.CategoryQuery
}

func NewCategoryQuery(categoryQuery domain.CategoryQuery) *CategoryQuery {
	return &CategoryQuery{
		categoryQuery: categoryQuery,
	}
}

func (c *CategoryQuery) FindAll() (*[]domain.Category, error) {
	categories, err := c.categoryQuery.FindAll()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *CategoryQuery) FindById(categoryId int64) (domain.Category, error) {
	category, err := c.categoryQuery.FindById(categoryId)
	return *category, err
}
