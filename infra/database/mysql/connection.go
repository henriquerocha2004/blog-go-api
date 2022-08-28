package mysql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type MysqlConnection struct {
	DbConnection *sql.DB
}
type MysqlConfig struct {
	User     string `json:"user"`
	PassWord string `json:"password"`
	Port     string `json:"port"`
	Database string `json:"db_name"`
	Host     string `json:"host"`
}

func NewMysqlConnection() *sql.DB {
	config := loadConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.PassWord,
		config.User,
		config.Host,
		config.Port,
		config.Database,
	)
	connection, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = connection.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	return connection
}

func loadConfig() *MysqlConfig {
	database, err := json.Marshal(viper.Get("database"))
	if err != nil {
		panic("failed to parse database config")
	}

	var mysqlConfig MysqlConfig
	err = json.Unmarshal(database, &mysqlConfig)
	if err != nil {
		panic("failed to parse database config")
	}
	return &mysqlConfig
}
