package command

import (
	domain "github.com/henriquerocha2004/blog-go-api/domain/entities"
	"github.com/henriquerocha2004/blog-go-api/infra/http/dto"
)

type UserAction struct {
	userCommand domain.UserCommand
}

func NewUserAction(userCommand domain.UserCommand) *UserAction {
	return &UserAction{
		userCommand: userCommand,
	}
}

func (u *UserAction) HandleCreate(userRequest dto.UserRequest) error {
	user := domain.User{
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Email:     userRequest.Email,
		PassWord:  userRequest.PassWord,
	}
	err := u.userCommand.Create(user)

	if err != nil {
		return err
	}
	return nil
}

func (u *UserAction) HandleUpdate(userId int64, userRequest dto.UserRequest) error {
	user := domain.User{
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Email:     userRequest.Email,
		PassWord:  userRequest.PassWord,
	}

	err := u.userCommand.Update(userId, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserAction) HandleDelete(userId int64) error {
	err := u.userCommand.Delete(userId)
	if err != nil {
		return err
	}
	return nil
}
