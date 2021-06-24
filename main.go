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
	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

type handler struct {
	db     *gorm.DB
	master *impl.Master
}

var master *impl.Master

var serviceImpl struct {
	getActivityProgramService impl.GetActivityProgramService
	putActivityProgramService impl.PutActivityProgramService
	getRoutesService          impl.GetRoutesService
	putRouteService           impl.PutRouteService
	deleteRouteService        impl.DeleteRouteService
	getHistories              impl.GetHistoriesService
	updateUsers               impl.UpdateUserService
	createUser                impl.CreateUserService
	getUser                   impl.GetUserService
	deleteActivityProgram     impl.DeleteActivityProgramService
	getActivityPrograms       impl.GetActivityProgramsService
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

	// dbConnection
	db, err := repository.NewDbConnection()
	if err != nil {
		log.Println(err)
		return
	}
	db.LogMode(true)
	defer db.Close()

	repo := repository.NewDbRepository(db)
	con, err := repo.NewConnection()
	if err != nil {
		log.Println(err)
		return
	}

	// set master data
	master, err = setMasterData(con)
	if err != nil {
		log.Println(err)
		return
	}

	h := &handler{
		db:     db,
		master: master,
	}

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("FRONT_URL")}
	config.AllowMethods = []string{"POST", "GET", "PUT", "DELETE"}
	config.AllowHeaders = []string{
		"Access-Control-Allow-Headers",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"X-CSRF-Token",
		"Authorization"}
	r.Use(cors.New(config))

	r.GET("/api/v1/get_activity_program/:seq_no", h.appHandler(&serviceImpl.getActivityProgramService))
	r.PUT("/api/v1/put_activity_program", h.appHandler(&serviceImpl.putActivityProgramService))
	r.GET("/api/v1/get_routes", h.appHandler(&serviceImpl.getRoutesService))
	r.PUT("/api/v1/put_route", h.appHandler(&serviceImpl.putRouteService))
	r.DELETE("/api/v1/delete_route", h.appHandler(&serviceImpl.deleteRouteService))
	r.GET("/api/v1/get_histories", h.appHandler(&serviceImpl.getHistories))
	r.PUT("/api/v1/update_user", h.appHandler(&serviceImpl.updateUsers))
	r.GET("/api/v1/get_user", h.appHandler(&serviceImpl.getUser))
	r.DELETE("/api/v1/delete_activity_program", h.appHandler(&serviceImpl.deleteActivityProgram))
	r.GET("/api/v1/get_activity_programs", h.appHandler(&serviceImpl.getActivityPrograms))

	if os.Getenv("NO_AUTH_FUNC_ON") == "1" {
		r.POST("/api/v1/create_user", h.appNoAuthHandler((&serviceImpl.createUser)))
	}

	r.Run()
}

func (h *handler) appHandler(s impl.ServiceImpl) func(*gin.Context) {
	return func(ctx *gin.Context) {
		req := s.New()

		repo := repository.NewDbRepository(h.db)
		con, err := repo.NewConnection()
		if err != nil {
			log.Println(err)
			errorHandring("db connections error", ctx)
			return
		}

		var token *auth.Token
		if os.Getenv("ENV") != "prd" {
			token, err = authDev(ctx)
			h.db.LogMode(true)
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
			errorHandring("servr error", ctx)
			return
		}

		implCtx := impl.NewContext(model.UserID(token.UID), con, master)
		res, err := req.Execute(implCtx)
		if err != nil {
			log.Println(err)
			errorHandring("servr error", ctx)
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}

func (h *handler) appNoAuthHandler(s impl.ServiceImpl) func(*gin.Context) {
	return func(ctx *gin.Context) {
		req := s.New()

		repo := repository.NewDbRepository(h.db)
		con, err := repo.NewConnection()
		if err != nil {
			log.Println(err)
			errorHandring("db connections error", ctx)
			return
		}

		req.SetRequest(ctx)
		err = req.Validate()
		if err != nil {
			log.Println(err)
			errorHandring("servr error", ctx)
			return
		}

		implCtx := impl.NewContext("", con, master)
		res, err := req.Execute(implCtx)
		if err != nil {
			log.Println(err)
			errorHandring("servr error", ctx)
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
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
		return nil, fmt.Errorf("failed Connection error")
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

func setMasterData(con repository.Connection) (*impl.Master, error) {
	practices := []*model.Practice{
		{ID: 1, Name: "レギュラー"},
		{ID: 2, Name: "卒演"},
		{ID: 3, Name: "新練"},
	}

	activities := []*model.Activity{
		{ID: 1, Name: "tutti"},
		{ID: 2, Name: "弦練"},
		{ID: 3, Name: "管打練"},
		{ID: 4, Name: "パート練"},
		{ID: 5, Name: "木管練"},
		{ID: 6, Name: "金管練"},
		{ID: 7, Name: "トップ練"},
		{ID: 8, Name: "引き継ぎ"},
		{ID: 9, Name: "アンサンブル"},
	}

	prices, err := con.ListPlace()
	if err != nil {
		return nil, err
	}
	return impl.NewMaster(practices, activities, prices), nil
}
