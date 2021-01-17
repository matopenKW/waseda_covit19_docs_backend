package impl

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

type HelloWorldImpl interface {
	HelloWorld(ctx *gin.Context) (*HelloWorldResponce, error)
}

type HelloWorldRequest struct {
	Name string
}

type HelloWorldResponce struct {
	Message string
}

func HelloWorld(con repository.Connection, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}
