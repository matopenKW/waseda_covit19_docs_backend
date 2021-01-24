package impl

import (
	"github.com/gin-gonic/gin"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

func UpdatePost(con repository.Connection, ctx *gin.Context) (ResponceImpl, error) {

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
