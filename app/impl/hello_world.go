package impl

import (
	"net/url"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

type HelloWorldRequest struct {
}

func (r *HelloWorldRequest) SetRequest(form url.Values) {

}

func (r *HelloWorldRequest) Validate() error {
	return nil
}

func (r *HelloWorldRequest) HelloWorld(con repository.Connection) (ResponceImpl, error) {
	return &HelloWorldResponce{
		Message: "hello world",
	}, nil
}
