package mariadb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

func ConnectWithConfig(config Config) (*sql.DB, error) {
	return Connect(config.Username, config.Password, config.Host, config.Port, config.DBName)
}

func Connect(host, port, dbName, username, password string) (*sql.DB, error) {
	connInfo := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		username, password, host, port, dbName,
	)
	db, err := sql.Open("mysql", connInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
