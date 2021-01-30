package impl

import "net/url"

type LoginRequest struct {
}

func (r *LoginRequest) SetRequest(form url.Values) {

}

func (r *LoginRequest) Validate() error {
	return nil
}

func (r *LoginRequest) Execute(ctx *Context) (ResponceImpl, error) {
	return login(r, ctx)
}

type LoginResponce struct {
	Message string
}

func (r *LoginResponce) GetResponce() {
}

func login(req *LoginRequest, ctx *Context) (ResponceImpl, error) {
	return &LoginResponce{
		Message: "hello world",
	}, nil
}
