package command

import (
	domain "github.com/henriquerocha2004/blog-go-api/domain/entities"
	"github.com/henriquerocha2004/blog-go-api/infra/http/dto"
)

type CategoryAction struct {
	categoryCommand domain.CategoryCommand
}

func NewCategoryAction(categoryCommand domain.CategoryCommand) *CategoryAction {
	return &CategoryAction{
		categoryCommand: categoryCommand,
	}
}

func (c *CategoryAction) HandleCreate(dtoCategory dto.Category) error {
	category := &domain.Category{
		Description: dtoCategory.Description,
	}

	err := c.categoryCommand.Create(*category)
	if err != nil {
		return err
	}

	return nil
}

func (c *CategoryAction) HandleUpdate(categoryId int64, dtoCategory dto.Category) error {
	category := &domain.Category{
		Description: dtoCategory.Description,
	}
	err := c.categoryCommand.Update(categoryId, *category)
	return err
}

func (c *CategoryAction) HandleDelete(categoryId int64) error {
	return c.categoryCommand.Delete(categoryId)
}
