package impl

import (
	"net/url"
)

type HelloWorldRequest struct {
}

func (r *HelloWorldRequest) SetRequest(form url.Values) {

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
