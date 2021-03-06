package routes

import "github.com/gin-gonic/gin"

var router = gin.Default()

func Run() {
	getRoutes()
	router.Run(":6000")

}
func getRoutes() {
	v1 := router.Group("v1")
	addUserRoutes(v1)
	addPingRoutes(v1)
	addAuthTotpRoutes(v1)

	v2 := router.Group("v2")
	addPingRoutes(v2)
	addAuthTotpRoutes(v2)
}
