package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "users")
	})

	users.GET("/comments", func(context *gin.Context) {
		context.JSON(http.StatusOK, "users commenets")
	})

	users.GET("/pictures", func(context *gin.Context) {
		context.JSON(http.StatusOK, "users pictures")
	})

}
