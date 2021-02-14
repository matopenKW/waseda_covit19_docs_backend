package impl

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
  )
  
  func TestHelloWorld_helloWorld(t *testing.T) {
	req, _ : = http.NewRequest("GET", "/test", nil)
  
	// Contextセット
	var context *gin.Context
	context = &gin.Context{Request: req}
  
	huga(context)
  }
  
  func huga(c *gin.Context) {
	....
  }