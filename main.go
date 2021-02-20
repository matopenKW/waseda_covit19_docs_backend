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

var serviceImpl struct {
	putActivityProgramService impl.PutActivityProgramService
	getRoutesService          impl.GetRoutesService
	putRouteService           impl.PutRouteService
	deleteRouteService        impl.DeleteRouteService
}

func init() {
	if os.Getenv("DATABASE_URL") == "" {
		panic("init error. db url env is brank")
	}
	if os.Getenv("FRONT_URL") == "" {
		panic("init error. front url env is brank")
	}
}

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("FRONT_URL")}
	config.AllowMethods = []string{"GET", "PUT", "DELETE"}
	config.AllowHeaders = []string{
		"Access-Control-Allow-Headers",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"X-CSRF-Token",
		"Authorization"}
	r.Use(cors.New(config))

	r.PUT("/api/v1/put_activity_program", appHandler(&serviceImpl.putActivityProgramService))
	r.GET("/api/v1/get_routes", appHandler(&serviceImpl.getRoutesService))
	r.PUT("/api/v1/put_route", appHandler(&serviceImpl.putRouteService))
	r.DELETE("/api/v1/delete_route", appHandler(&serviceImpl.deleteRouteService))

	r.Run()
}

func appHandler(s impl.ServiceImpl) func(*gin.Context) {
	return func(ctx *gin.Context) {
		req := s.New()

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

		req.SetRequest(ctx)
		err = req.Validate()
		if err != nil {
			log.Println(err)
			errorHandring("servr ereror", ctx)
			return
		}

		implCtx := impl.NewContext(token.UID, con)
		res, err := req.Execute(implCtx)
		if err != nil {
			log.Println(err)
			errorHandring("servr ereror", ctx)
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
		"message": message,
	})
}

func authJWT(ctx *gin.Context) (*auth.Token, error) {
	auth, err := repository.OpenAuthJSON()
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
