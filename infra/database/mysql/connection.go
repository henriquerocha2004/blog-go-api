package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlConnection struct {
	DbConnection *sql.DB
}

func NewMysqlConnection() *sql.DB {
	connection, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/blogapi")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = connection.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	return connection
}
