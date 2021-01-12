package impl

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(db *sql.DB, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}
