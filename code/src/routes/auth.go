package routes

import (
	"github.com/gin-gonic/gin"
	auth "google-authenticator/src/authenticator"
	"net/http"
)

func addAuthTotpRoutes(rg *gin.RouterGroup) {
	totps := rg.Group("/totp")

	totps.GET("/id", func(context *gin.Context) {
		context.JSON(http.StatusOK, "")
	})

	totps.POST("/add", func(context *gin.Context) {
		context.JSON(http.StatusOK, "")
	})

	totps.DELETE("/id", func(context *gin.Context) {
		context.JSON(http.StatusOK, "")
	})

	totps.POST("/update", func(context *gin.Context) {
		context.JSON(http.StatusOK, "")
	})

	totps.GET("/list", func(context *gin.Context) {
		var passwords []auth.OneTimePassword

		for totp := range auth.OtpMap {
			passwd, _ := auth.GetOneTimePassword(auth.OtpMap[totp].Secret, auth.OtpMap[totp].Name)
			if passwd != nil {
				passwords = append(passwords, *passwd)
			}
		}
		context.JSON(http.StatusOK, gin.H{
			"data": passwords,
		})
	})

	totps.GET("/:id", func(context *gin.Context) {
		tpid := context.Param("id")
		opt := auth.OtpMap[tpid]
		passwd, _ := auth.GetOneTimePassword(opt.Secret, opt.Name)
		context.JSON(http.StatusOK, gin.H{
			"data": *passwd,
		})
	})
}

func addAuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")

	auth.POST("/login", func(context *gin.Context) {

		//{"authorized":true,"id":65,"token":"52465bfb-10da-4ee9-b887-3bd37ccb297a","username":"jiangtao","displayName":"蒋涛","avator":"1.png","orgId":1,"domain":"anxinyun","orgName":"江西飞尚","orgtype":1001,"portal":"A","departmentId":1,"departmentName":"默认","resources":[""]}

		context.JSON(http.StatusOK, "")
	})

}
