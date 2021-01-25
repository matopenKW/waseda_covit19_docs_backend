package impl

import (
	"net/url"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

type PutPostRequest struct {
}

func (r *PutPostRequest) SetRequest(form url.Values) {

}

func (r *PutPostRequest) Validate() error {
	return nil
}

func (r *PutPostRequest) Execute(con repository.Connection) (ResponceImpl, error) {
	return putPost(r, con)
}

type PutPostResponce struct {
	Post *model.Post
}

func (r *PutPostResponce) GetResponce() {
}

func putPost(req *PutPostRequest, con repository.Connection) (ResponceImpl, error) {
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
