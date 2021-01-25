package impl

import (
	"net/url"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

type GetPostsRequest struct {
}

func (r *GetPostsRequest) SetRequest(form url.Values) {

}

func (r *GetPostsRequest) Validate() error {
	return nil
}

func (r *GetPostsRequest) Execute(con repository.Connection) (ResponceImpl, error) {
	return GetPosts(r, con)
}

type GetPostsResponce struct {
	Posts []*model.Post
}

func (r *GetPostsResponce) GetResponce() {
}

func GetPosts(req *GetPostsRequest, con repository.Connection) (ResponceImpl, error) {
	ps, err := con.GetPosts()
	if err != nil {
		return nil, err
	}

	return &GetPostsResponce{
		Posts: ps,
	}, nil
}
