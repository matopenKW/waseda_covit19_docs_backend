package impl

import "github.com/gin-gonic/gin"

type HelloWorldRequest struct {
}

func (r *HelloWorldRequest) SetRequest(ctx *gin.Context) {

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
		Message: "hello world",
	}, nil
}
