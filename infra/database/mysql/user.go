package mysql

import (
	"database/sql"
	"log"

	domain "github.com/henriquerocha2004/blog-go-api/domain/entities"
)

type UserCommand struct {
	connection *sql.DB
}

type UserQuery struct {
	connection *sql.DB
}

func NewUserCommand(connection *sql.DB) *UserCommand {
	return &UserCommand{
		connection: connection,
	}
}

func (userCmd *UserCommand) Create(user domain.User) error {
	hashPassword, err := user.HashPassword()
	if err != nil {
		return err
	}

	stmt, err := userCmd.connection.Prepare("INSERT INTO users (first_name, last_name, email, password) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, hashPassword)
	if err != nil {
		return err
	}

	return nil
}

func (userCmd *UserCommand) Update(userId int64, user domain.User) error {
	stmt, err := userCmd.connection.Prepare("UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, userId)
	if err != nil {
		return err
	}

	return nil
}

func (userCmd *UserCommand) Delete(userId int64) error {

	stmt, err := userCmd.connection.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userId)
	if err != nil {
		return err
	}

	return nil
}

func NewUserQuery(connection *sql.DB) *UserQuery {
	return &UserQuery{
		connection: connection,
	}
}

func (usrQry *UserQuery) FindById(userId int64) (domain.User, error) {

	rows, err := usrQry.connection.Query("SELECT id, first_name, last_name, email FROM users WHERE id = ?", userId)
	if err != nil {
		return domain.User{}, err
	}

	defer rows.Close()
	var user domain.User

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
	}
	return user, nil
}

func (usrQry *UserQuery) FindAll() (*[]domain.User, error) {
	rows, err := usrQry.connection.Query("SELECT id, first_name, last_name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []domain.User

	for rows.Next() {
		var user domain.User
		rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email)
		users = append(users, user)
	}
	return &users, nil
}
