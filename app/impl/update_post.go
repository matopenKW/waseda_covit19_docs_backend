package impl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func UpdatePost(db *gorm.DB, c *gin.Context) {
	result := db.Save(&Post{
		ID:      1,
		Content: "content1-1",
		Author:  "author1-1",
	})
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "dbError",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": result.Value,
	})
}
