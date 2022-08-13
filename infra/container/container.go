package container

import (
	"database/sql"

	"github.com/henriquerocha2004/blog-go-api/domain/actions/command"
	"github.com/henriquerocha2004/blog-go-api/domain/actions/query"
	domain "github.com/henriquerocha2004/blog-go-api/domain/entities"
	"github.com/henriquerocha2004/blog-go-api/infra/database/mysql"
	"github.com/henriquerocha2004/blog-go-api/infra/http/controllers"
)

type ContainerDependency struct {
	connection *sql.DB

	userCommand domain.UserCommand
	userQuery   domain.UserQuery

	userController *controllers.UserController

	userHandleAction *command.UserAction
	userHandleQuery  *query.UserQuery
}

func (c *ContainerDependency) GetDatabaseConnection() *sql.DB {
	if c.connection == nil {
		c.connection = mysql.NewMysqlConnection()
	}
	return c.connection
}

func (c *ContainerDependency) GetUserCommand() domain.UserCommand {
	if c.userCommand == nil {
		c.userCommand = mysql.NewUserCommand(
			c.GetDatabaseConnection(),
		)
	}
	return c.userCommand
}

func (c *ContainerDependency) GetUserQuery() domain.UserQuery {
	if c.userQuery == nil {
		c.userQuery = mysql.NewUserQuery(
			c.GetDatabaseConnection(),
		)
	}

	return c.userQuery
}

func (c *ContainerDependency) GetUserHandleAction() *command.UserAction {
	c.userHandleAction = command.NewUserAction(
		c.GetUserCommand(),
	)
	return c.userHandleAction
}

func (c *ContainerDependency) GetUserHandleQuery() *query.UserQuery {
	c.userHandleQuery = query.NewUserQuery(
		c.GetUserQuery(),
	)
	return c.userHandleQuery
}

func (c *ContainerDependency) GetUserController() *controllers.UserController {
	c.userController = controllers.NewUserController(
		c.GetUserHandleAction(),
		c.GetUserHandleQuery(),
	)
	return c.userController
}
