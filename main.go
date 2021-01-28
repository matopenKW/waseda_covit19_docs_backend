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

func init() {
	_, err := dbConnection()
	if err != nil {
		fmt.Println(err.Error())
		panic("error db connection")
	}
}

func main() {
	r := gin.Default()

	r.GET("/api/v1/hello_world", appHandler(&impl.HelloWorldRequest{}))
	r.GET("/api/v1/post", appHandler(&impl.GetPostsRequest{}))
	r.GET("/api/v1/post_put", appHandler(&impl.PutPostRequest{}))
	r.GET("/api/v1/post_update", appHandler(&impl.UpdatePostRequest{}))
	r.GET("/api/v1/get_routes", appHandler(&impl.GetRoutesRequest{}))

	r.Run()
}

func appHandler(i impl.RequestImpl) func(*gin.Context) {
	return func(ctx *gin.Context) {
		// dbConnection
		db, err := dbConnection()
		defer db.Close()
		if err != nil {
			errorHandring("db connections error", ctx)
			return
		}

		repo := repository.NewDbRepository(db)
		con, err := repo.NewConnection()
		if err != nil {
			errorHandring("db connections error", ctx)
			return
		}

		req := ctx.Request
		req.ParseForm()
		i.SetRequest(req.Form)
		i.Validate()

		implCtx := impl.NewContext("user_id", con)
		res, err := i.Execute(implCtx)
		if err != nil {
			errorHandring("server error", ctx)
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}

func dbConnection() (*gorm.DB, error) {
	return gorm.Open("postgres", os.Getenv("DATABASE_URL"))
}

func errorHandring(message string, ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": message,
	})
}
