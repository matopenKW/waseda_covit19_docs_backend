package impl

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

func GetPosts(con repository.Connection, c *gin.Context) {
	ps, err := con.GetPosts()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "dbError",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": ps,
	})
}
