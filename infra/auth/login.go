package auth

import (
	"errors"
	domain "github.com/henriquerocha2004/blog-go-api/domain/entities"
	"github.com/henriquerocha2004/blog-go-api/infra/http/dto"
)

type Login struct {
	userQuery domain.UserQuery
}

type TokenResponse struct {
	Token string `json:"token"`
}

func NewLogin(userQuery domain.UserQuery) *Login {
	return &Login{
		userQuery: userQuery,
	}
}

func (l *Login) Authenticate(userCredentials dto.UserCredentials) (*TokenResponse, error) {
	user, err := l.userQuery.FindByEmail(userCredentials.Email)
	if err != nil || user.Id == 0 {
		return nil, errors.New("invalid credentials")
	}

	if err := user.CheckPassword(userCredentials.Password); err != nil {
		return nil, errors.New("invalid credentials")
	}
	token, err := GenerateToken(user.Email, user.FirstName, user.LastName)
	var tokenResponse TokenResponse
	tokenResponse.Token = token
	return &tokenResponse, err
}
