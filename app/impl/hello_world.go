package impl

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type HelloWorldRequest struct {
	Message string `json:"message"`
}

func (r *HelloWorldRequest) SetRequest(ctx *gin.Context) {
	_ = ctx.ShouldBindJSON(&r)
}

func (r *HelloWorldRequest) Validate() error {
	return nil
}

func (r *HelloWorldRequest) Execute(ctx *Context) (ResponceImpl, error) {
	return helloWorld(r, ctx)
}

type HelloWorldResponce struct {
	Message string
}

func (r *HelloWorldResponce) GetResponce() {
}

func helloWorld(req *HelloWorldRequest, ctx *Context) (ResponceImpl, error) {
	return &HelloWorldResponce{
		Message: fmt.Sprintf("hello world %s", req.Message),
	}, nil
}
