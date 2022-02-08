package mariadb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(username, password, host, port, dbName string) (*sql.DB, error) {
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
