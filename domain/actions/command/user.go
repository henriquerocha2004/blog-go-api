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

}
