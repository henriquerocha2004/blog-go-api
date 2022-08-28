package mysql_test

import (
	"testing"

	domain "github.com/henriquerocha2004/blog-go-api/domain/entities"
	"github.com/henriquerocha2004/blog-go-api/infra/mocks"
	"github.com/stretchr/testify/suite"
)

type TestUserSuit struct {
	suite.Suite
	commandUser domain.UserCommand
	queryUser   domain.UserQuery
}

func newTestUserSuit() *TestUserSuit {
	return &TestUserSuit{}
}

func TestUserTests(t *testing.T) {
	suite.Run(t, newTestUserSuit())
}

func (s *TestUserSuit) TestCreateUser() {
	user := domain.User{
		FirstName: "Henrique",
		LastName:  "Souza",
		Email:     "rochahenrique18@gmail.com",
		PassWord:  "Teste123",
	}

	mockCommand := new(mocks.UserCommand)
	mockCommand.On("Create", user).Return(nil)
	s.commandUser = mockCommand
	err := s.commandUser.Create(user)
	s.NoError(err)
}

func (s *TestUserSuit) TestUpdateUser() {
	user := domain.User{
		FirstName: "José",
		LastName:  "Souza",
		Email:     "rochahenrique18@gmail.com",
		PassWord:  "Teste123",
	}
	mockCommand := new(mocks.UserCommand)
	mockCommand.On("Update", int64(1), user).Return(nil)

	s.commandUser = mockCommand
	err := s.commandUser.Update(1, user)
	s.NoError(err)
}

func (s *TestUserSuit) TestDeleteUser() {
	mockCommand := new(mocks.UserCommand)
	mockCommand.On("Delete", int64(1)).Return(nil)
	s.commandUser = mockCommand
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

	mockQuery := new(mocks.UserQuery)
	mockQuery.On("FindById", int64(1)).Return(user, nil)
	s.queryUser = mockQuery
	userDb, err := s.queryUser.FindById(int64(1))

	s.NoError(err)
	s.Equal(user.FirstName, userDb.FirstName)
	s.Equal(user.LastName, userDb.LastName)
	s.Equal(user.Email, userDb.Email)
}

func (s *TestUserSuit) TestFindAll() {
	usersMock := []domain.User{
		{
			FirstName: "Henrique",
			LastName:  "Souza",
			Email:     "rochahenrique18@gmail.com",
			PassWord:  "Teste123",
		},
		{
			FirstName: "Luciana",
			LastName:  "Souza",
			Email:     "luciana@gmail.com",
			PassWord:  "Teste123",
		},
	}

	mockQuery := new(mocks.UserQuery)
	mockQuery.On("FindAll").Return(&usersMock, nil)
	s.queryUser = mockQuery

	users, err := s.queryUser.FindAll()
	s.NoError(err)
	s.Equal(2, len(*users))
}
