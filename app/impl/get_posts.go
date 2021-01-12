package impl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Post struct {
	ID      int
	Content string
	Author  string
}

func GetPosts(db *gorm.DB, c *gin.Context) {
	var ps []*Post
	err := db.Find(&ps).Error
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
