package api

import (
	"github.com/crissilvaeng/goseed/api/routes"

	"github.com/gin-gonic/gin"
)

// Run define the handled endpoint on API.
func Run(port string) {
	r := gin.Default()

	v1 := r.Group("/v1")

	v1.GET("/hello", routes.Hello)

	r.Run(":" + port)
}
