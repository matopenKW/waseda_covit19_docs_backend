package impl

import (
	"github.com/gin-gonic/gin"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

func PutPost(con repository.Connection, ctx *gin.Context) (ResponceImpl, error) {
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
