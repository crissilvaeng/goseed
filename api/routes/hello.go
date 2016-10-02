package routes

import (
	"fmt"
	"net/http"

	"github.com/crissilvaeng/goddamned/api/universe"
	"github.com/gin-gonic/gin"
)

// Hello provides a endpoint for /hello request.
func Hello(c *gin.Context) {

	answer := fmt.Sprintf("The answer to life, the universe and everything is %d.", universe.Answer())

	c.JSON(http.StatusOK, gin.H{
		"message": answer,
	})
}
