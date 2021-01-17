package impl

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

func PutPost(con repository.Connection, ctx *gin.Context) {
	p := &model.Post{
		ID:      2,
		Content: "content2",
		Author:  "author2",
	}
	result, err := con.CreatePost(p)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "dbError",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"post": result,
	})
}
