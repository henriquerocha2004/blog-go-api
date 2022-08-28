package container

import (
	"database/sql"
	"github.com/henriquerocha2004/blog-go-api/infra/auth"

	"github.com/henriquerocha2004/blog-go-api/domain/actions/command"
	"github.com/henriquerocha2004/blog-go-api/domain/actions/query"
	domain "github.com/henriquerocha2004/blog-go-api/domain/entities"
	"github.com/henriquerocha2004/blog-go-api/infra/database/mysql"
	"github.com/henriquerocha2004/blog-go-api/infra/http/controllers"
)

type ContainerDependency struct {
	connection *sql.DB

	userCommand     domain.UserCommand
	userQuery       domain.UserQuery
	categoryCommand domain.CategoryCommand
	categoryQuery   domain.CategoryQuery
	postCommand     domain.PostCommand
	postQuery       domain.PostQuery

	userController         *controllers.UserController
	categoryController     *controllers.CategoryController
	postController         *controllers.PostController
	authenticateController *controllers.AuthController

	userHandleAction     *command.UserAction
	userHandleQuery      *query.UserQuery
	categoryHandleAction *command.CategoryAction
	categoryHandleQuery  *query.CategoryQuery
	postHandlerCommand   *command.PostAction
	postHandlerQuery     *query.PostQuery

	login *auth.Login
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

func (c *ContainerDependency) GetCategoryCommand() domain.CategoryCommand {
	if c.categoryCommand == nil {
		c.categoryCommand = mysql.NewCategoryCommand(
			c.GetDatabaseConnection(),
		)
	}

	return c.categoryCommand
}

func (c *ContainerDependency) GetCategoryQuery() domain.CategoryQuery {
	if c.categoryQuery == nil {
		c.categoryQuery = mysql.NewCategoryQuery(
			c.GetDatabaseConnection(),
		)
	}

	return c.categoryQuery
}

func (c *ContainerDependency) GetPostCommand() domain.PostCommand {
	if c.postCommand == nil {
		c.postCommand = mysql.NewPostCommand(
			c.GetDatabaseConnection(),
		)
	}
	return c.postCommand
}

func (c *ContainerDependency) GetPostQuery() domain.PostQuery {
	if c.postQuery == nil {
		c.postQuery = mysql.NewPostQuery(
			c.GetDatabaseConnection(),
		)
	}
	return c.postQuery
}

func (c *ContainerDependency) GetPostHandlerQuery() *query.PostQuery {
	c.postHandlerQuery = query.NewPostQuery(
		c.GetPostQuery(),
	)
	return c.postHandlerQuery
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

func (c *ContainerDependency) GetCategoryAction() *command.CategoryAction {
	c.categoryHandleAction = command.NewCategoryAction(
		c.GetCategoryCommand(),
	)
	return c.categoryHandleAction
}

func (c *ContainerDependency) GetCategoryHandlerQuery() *query.CategoryQuery {
	c.categoryHandleQuery = query.NewCategoryQuery(
		c.GetCategoryQuery(),
	)

	return c.categoryHandleQuery
}

func (c *ContainerDependency) GetPostHandlerCommand() *command.PostAction {
	c.postHandlerCommand = command.NewPostAction(
		c.GetPostCommand(),
	)
	return c.postHandlerCommand
}

func (c *ContainerDependency) GetUserController() *controllers.UserController {
	c.userController = controllers.NewUserController(
		c.GetUserHandleAction(),
		c.GetUserHandleQuery(),
	)
	return c.userController
}

func (c *ContainerDependency) GetCategoryController() *controllers.CategoryController {
	if c.categoryController == nil {
		c.categoryController = controllers.NewCategoryController(
			c.GetCategoryAction(),
			c.GetCategoryHandlerQuery(),
		)
	}
	return c.categoryController
}

func (c *ContainerDependency) GetPostController() *controllers.PostController {
	if c.postController == nil {
		c.postController = controllers.NewPostController(
			c.GetPostHandlerCommand(),
			c.GetPostHandlerQuery(),
		)
	}
	return c.postController
}

func (c *ContainerDependency) GetAuthenticateController() *controllers.AuthController {
	if c.authenticateController == nil {
		c.authenticateController = controllers.NewAuthController(
			c.GetActionLogin(),
		)
	}
	return c.authenticateController
}

func (c *ContainerDependency) GetActionLogin() *auth.Login {
	c.login = auth.NewLogin(
		c.GetUserQuery(),
	)
	return c.login
}
