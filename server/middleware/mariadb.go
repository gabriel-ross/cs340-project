package middleware

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Recover(db *sql.DB, connStr string) gin.HandlerFunc {
	return func(c *gin.Context) {

		err := db.Ping()
		if err != nil {
			//
		}

		c.Next()
	}
}
