package impl

import (
	"github.com/gin-gonic/gin"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

func GetPosts(con repository.Connection, c *gin.Context) (ResponceImpl, error) {
	ps, err := con.GetPosts()
	if err != nil {
		return nil, err
	}

	return &GetPostsResponce{
		Posts: ps,
	}, nil
}
