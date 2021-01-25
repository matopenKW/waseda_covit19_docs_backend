package impl

import (
	"net/url"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

type UpdatePostRequest struct {
}

func (r *UpdatePostRequest) SetRequest(form url.Values) {

}

func (r *UpdatePostRequest) Validate() error {
	return nil
}

func (r *UpdatePostRequest) Execute(con repository.Connection) (ResponceImpl, error) {
	return UpdatePost(r, con)
}

type UpdatePostResponce struct {
	Post *model.Post
}

func (r *UpdatePostResponce) GetResponce() {
}

func UpdatePost(req *UpdatePostRequest, con repository.Connection) (ResponceImpl, error) {

	result, err := con.SavePost(&model.Post{
		ID:      1,
		Content: "content1-1",
		Author:  "author1-1",
	})
	if err != nil {
		return nil, err
	}
	return &UpdatePostResponce{
		Post: result,
	}, nil
}
