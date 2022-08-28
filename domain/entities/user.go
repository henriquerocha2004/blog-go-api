package domain

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        int64  `json:"id,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	PassWord  string `json:"password,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

func (u *User) HashPassword() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.PassWord), 14)
	return string(bytes), err
}

func (u *User) CheckPassword(providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.PassWord), []byte(providedPassword))
}

type UserCommand interface {
	Create(user User) error
	Update(userId int64, user User) error
	Delete(userId int64) error
}

type UserQuery interface {
	FindById(userId int64) (User, error)
	FindByEmail(email string) (User, error)
	FindAll() (*[]User, error)
}
