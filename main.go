package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/impl"
)

type HelloImpl interface {
	HelloWorld(c *gin.Context)
}

func init() {
	_, err := dbConnection()
	if err != nil {
		fmt.Println(err.Error())
		panic("error db connection")
	}
}

func main() {
	r := gin.Default()

	r.GET("/api/v1/hello_world", appHandler(impl.HelloWorld))
	r.GET("/bye", appHandler(impl.HelloWorld))

	r.Run()
}

func appHandler(i func(*sql.DB, *gin.Context)) func(*gin.Context) {
	return func(c *gin.Context) {
		// dbConnection
		db, err := dbConnection()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "db connections error",
			})
			return
		}
		i(db, c)
	}
}

func dbConnection() (*sql.DB, error) {
	return sql.Open("postgres", os.Getenv("DATABASE_URL"))
}
