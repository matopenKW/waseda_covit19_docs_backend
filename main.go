package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/gin-contrib/cors"
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

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	r.GET("/api/v1/get_routes", appHandler(&impl.GetRoutesRequest{}))
	r.PUT("/api/v1/put_route", appHandler(&impl.PutRouteRequest{}))
	r.DELETE("/api/v1/delete_route", appHandler(&impl.PutRouteRequest{}))

	r.Run()
}

func appHandler(i impl.RequestImpl) func(*gin.Context) {
	return func(ctx *gin.Context) {
		// dbConnection
		db, err := dbConnection()
		defer db.Close()
		if err != nil {
			log.Println(err)
			errorHandring("db connections error", ctx)
			return
		}
		db.LogMode(true)

		repo := repository.NewDbRepository(db)
		con, err := repo.NewConnection()
		if err != nil {
			log.Println(err)
			errorHandring("db connections error", ctx)
			return
		}

		var token *auth.Token
		if os.Getenv("ENV") != "prd" {
			token, err = authDev(ctx)
		} else {
			token, err = authJWT(ctx)
		}
		if err != nil {
			log.Println(err)
			errorHandring(err.Error(), ctx)
			return
		}

		i.SetRequest(ctx)
		i.Validate()

		implCtx := impl.NewContext(token.UID, con)
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

func authJWT(ctx *gin.Context) (*auth.Token, error) {
	auth, err := repository.OpenAuth()
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("Failed Connection error")
	}

	authHeader := ctx.Request.Header.Get("Authorization")
	idToken := strings.Replace(authHeader, "Bearer ", "", 1)

	token, err := auth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(fmt.Sprintf("error verifying ID token: %v\n", err))
	}

	return token, nil
}

func authDev(ctx *gin.Context) (*auth.Token, error) {
	return &auth.Token{
		UID: "user_id",
	}, nil
}
