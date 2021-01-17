package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/impl"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
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
	r.GET("/api/v1/post", appHandler(impl.GetPosts))
	r.GET("/api/v1/post_put", appHandler(impl.PutPost))
	r.GET("/api/v1/post_update", appHandler(impl.UpdatePost))
	r.GET("/bye", appHandler(impl.HelloWorld))

	r.Run()
}

func appHandler(i func(repository.Connection, *gin.Context)) func(*gin.Context) {
	return func(c *gin.Context) {
		// dbConnection
		db, err := dbConnection()
		defer db.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "db connections error",
			})
			return
		}

		repo := repository.NewDbRepository(db)
		con, err := repo.NewConnection()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "db connections error",
			})
			return
		}

		i(con, c)
	}
}

func dbConnection() (*gorm.DB, error) {
	return gorm.Open("postgres", os.Getenv("DATABASE_URL"))
}
