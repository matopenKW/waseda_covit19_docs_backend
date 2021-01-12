package impl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func PutPost(db *gorm.DB, c *gin.Context) {
	result := db.Create(&Post{
		ID:      2,
		Content: "content2",
		Author:  "author2",
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
