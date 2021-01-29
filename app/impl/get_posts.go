package impl

import (
	"net/url"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
)

type GetPostsRequest struct {
}

func (r *GetPostsRequest) SetRequest(form url.Values) {

}

func (r *GetPostsRequest) Validate() error {
	return nil
}

func (r *GetPostsRequest) Execute(ctx *Context) (ResponceImpl, error) {
	return getPosts(r, ctx)
}

type GetPostsResponce struct {
	Posts []*model.Post
}

func (r *GetPostsResponce) GetResponce() {
}

func getPosts(req *GetPostsRequest, ctx *Context) (ResponceImpl, error) {
	con := ctx.GetConnection()
	ps, err := con.GetPosts()
	if err != nil {
		return nil, err
	}

	return &GetPostsResponce{
		Posts: ps,
	}, nil
}
