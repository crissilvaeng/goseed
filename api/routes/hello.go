package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Hello provides a endpoint for /hello request.
func Hello(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
