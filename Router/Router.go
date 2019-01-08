package Router

import (
	"TimeLine/App"
	"TimeLine/Lib"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	Prefix_Api := router.Group("/api")

	//API Version
	V1 := Prefix_Api.Group("/v1")

	//auth
	Auth_V1 := V1.Group("/", Lib.JWTAuth())
	//Auth_V1.GET("/hello", Api.HelloServer)
	Auth_V1.GET("/growthstandards/:skip/:limit", Api.GetGrowthStandards)
	V1.POST("/person/add", Api.CreatePerson)

	//ignore auth
	NotAuth_V1 := V1.Group("/")
	NotAuth_V1.POST("/login", Api.Login)
	NotAuth_V1.POST("/wxlogin", Api.Wechat)

}
