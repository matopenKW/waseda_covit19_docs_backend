package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
)

type UpdatePostRequest struct {
}

func (r *UpdatePostRequest) SetRequest(ctx *gin.Context) {

}

func (r *UpdatePostRequest) Validate() error {
	return nil
}

func (r *UpdatePostRequest) Execute(ctx *Context) (ResponceImpl, error) {
	return UpdatePost(r, ctx)
}

type UpdatePostResponce struct {
	Post *model.Post
}

func (r *UpdatePostResponce) GetResponce() {
}

func UpdatePost(req *UpdatePostRequest, ctx *Context) (ResponceImpl, error) {
	con := ctx.GetConnection()

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
