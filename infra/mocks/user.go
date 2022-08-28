package mocks

import (
	domain "github.com/henriquerocha2004/blog-go-api/domain/entities"
	"github.com/stretchr/testify/mock"
)

type UserCommand struct {
	mock.Mock
}

func (u *UserCommand) Create(user domain.User) error {
	args := u.Called(user)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (u *UserCommand) Update(userId int64, user domain.User) error {
	args := u.Called(userId, user)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (u *UserCommand) Delete(userId int64) error {
	args := u.Called(userId)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

type UserQuery struct {
	mock.Mock
}

func (u *UserQuery) FindByEmail(email string) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserQuery) FindAll() (*[]domain.User, error) {
	args := u.Called()

	if err, ok := args.Get(1).(error); ok {
		return nil, err
	}
	return args.Get(0).(*[]domain.User), nil
}

func (u *UserQuery) FindById(userId int64) (domain.User, error) {
	args := u.Called(userId)

	if err, ok := args.Get(1).(error); ok {
		return domain.User{}, err
	}

	return args.Get(0).(domain.User), nil
}
