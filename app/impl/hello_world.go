package impl

import (
	"github.com/gin-gonic/gin"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

type HelloWorldRequest struct {
	Name string
}

func HelloWorld(con repository.Connection, c *gin.Context) (ResponceImpl, error) {
	return &HelloWorldResponce{
		Message: "hello world",
	}, nil
}
