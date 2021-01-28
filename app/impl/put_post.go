package impl

import (
	"net/url"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
)

type PutPostRequest struct {
}

func (r *PutPostRequest) SetRequest(form url.Values) {

}

func (r *PutPostRequest) Validate() error {
	return nil
}

func (r *PutPostRequest) Execute(ctx *Context) (ResponceImpl, error) {
	return putPost(r, ctx)
}

type PutPostResponce struct {
	Post *model.Post
}

func (r *PutPostResponce) GetResponce() {
}

func putPost(req *PutPostRequest, ctx *Context) (ResponceImpl, error) {
	con := ctx.GetConnection()
	p := &model.Post{
		ID:      2,
		Content: "content2",
		Author:  "author2",
	}
	result, err := con.CreatePost(p)
	if err != nil {
		return nil, err
	}

	return &PutPostResponce{
		Post: result,
	}, nil
}
