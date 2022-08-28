package mysql

import (
	"database/sql"

	domain "github.com/henriquerocha2004/blog-go-api/domain/entities"
)

type CategoryCommand struct {
	connection *sql.DB
}

func NewCategoryCommand(connection *sql.DB) *CategoryCommand {
	return &CategoryCommand{
		connection: connection,
	}
}

func (c *CategoryCommand) Create(category domain.Category) error {

	stmt, err := c.connection.Prepare("INSERT INTO categories (description) VALUES (?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(category.Description)
	if err != nil {
		return err
	}

	return nil
}

func (c *CategoryCommand) Delete(categoryId int64) error {
	stmt, err := c.connection.Prepare("DELETE FROM categories WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(categoryId)
	if err != nil {
		return err
	}

	return nil
}

func (c *CategoryCommand) Update(categoryId int64, category domain.Category) error {
	stmt, err := c.connection.Prepare("UPDATE categories SET description = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(category.Description, categoryId)
	if err != nil {
		return err
	}

	return nil
}

type CategoryQuery struct {
	connection *sql.DB
}

func NewCategoryQuery(connection *sql.DB) *CategoryQuery {
	return &CategoryQuery{
		connection: connection,
	}
}

func (c *CategoryQuery) FindAll() (*[]domain.Category, error) {
	rows, err := c.connection.Query("SELECT id, description FROM categories")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []domain.Category

	for rows.Next() {
		var category domain.Category
		err := rows.Scan(&category.Id, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return &categories, nil
}

func (c *CategoryQuery) FindById(categoryId int64) (*domain.Category, error) {
	row, err := c.connection.Query("SELECT id, description FROM categories WHERE id = ?", categoryId)
	if err != nil {
		return nil, err
	}

	defer row.Close()
	var category domain.Category

	for row.Next() {
		err = row.Scan(&category.Id, &category.Description)
		if err != nil {
			return nil, err
		}
	}
	return &category, nil
}
