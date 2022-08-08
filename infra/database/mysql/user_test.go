package mysql_test

import (
	"database/sql"
	"testing"

	domain "github.com/henriquerocha2004/blog-go-api/domain/entities"
	"github.com/henriquerocha2004/blog-go-api/infra/database/mysql"
	"github.com/stretchr/testify/suite"
)

type TestUserSuit struct {
	suite.Suite
	connection  *sql.DB
	commandUser domain.UserCommand
	queryUser   domain.UserQuery
}

func newTestUserSuit() *TestUserSuit {
	return &TestUserSuit{
		connection: mysql.NewMysqlConnection(),
	}
}

func TestUserTests(t *testing.T) {
	suite.Run(t, newTestUserSuit())
}

func (s *TestUserSuit) SetupSuite() {
	s.commandUser = mysql.NewUserCommand(s.connection)
	s.queryUser = mysql.NewUserQuery(s.connection)
}

func (suite *TestUserSuit) BeforeTest(suiteName, testName string) {
	suite.connection.Query("TRUNCATE TABLE users")
}

func (s *TestUserSuit) TestCreateUser() {
	userCommand := mysql.NewUserCommand(s.connection)
	user := domain.User{
		FirstName: "Henrique",
		LastName:  "Souza",
		Email:     "rochahenrique18@gmail.com",
		PassWord:  "Teste123",
	}
	err := userCommand.Create(user)
	s.NoError(err)
}

func (s *TestUserSuit) TestUpdateUser() {
	user := domain.User{
		FirstName: "Henrique",
		LastName:  "Souza",
		Email:     "rochahenrique18@gmail.com",
		PassWord:  "Teste123",
	}

	s.commandUser.Create(user)

	user.FirstName = "José"
	err := s.commandUser.Update(1, user)
	s.NoError(err)
}

func (s *TestUserSuit) TestDeleteUser() {
	user := domain.User{
		FirstName: "Henrique",
		LastName:  "Souza",
		Email:     "rochahenrique18@gmail.com",
		PassWord:  "Teste123",
	}

	s.commandUser.Create(user)
	err := s.commandUser.Delete(1)
	s.NoError(err)
}

func (s *TestUserSuit) TestFindUserById() {
	user := domain.User{
		FirstName: "Henrique",
		LastName:  "Souza",
		Email:     "rochahenrique18@gmail.com",
		PassWord:  "Teste123",
	}

	s.commandUser.Create(user)
	userDb, err := s.queryUser.FindById(1)

	s.NoError(err)
	s.Equal(user.FirstName, userDb.FirstName)
	s.Equal(user.LastName, userDb.LastName)
	s.Equal(user.Email, userDb.Email)
}

func (s *TestUserSuit) TestFindAll() {
	user := domain.User{
		FirstName: "Henrique",
		LastName:  "Souza",
		Email:     "rochahenrique18@gmail.com",
		PassWord:  "Teste123",
	}

	s.commandUser.Create(user)

	user = domain.User{
		FirstName: "Luciana",
		LastName:  "Souza",
		Email:     "luciana@gmail.com",
		PassWord:  "Teste123",
	}

	s.commandUser.Create(user)
	users, err := s.queryUser.FindAll()

	s.NoError(err)
	s.Equal(2, len(*users))
}
