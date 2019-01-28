package Router

import (
	ApiRoute "TimeLine/App"
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
	//Auth_V1.GET("/hello", ApiRoute.HelloServer)
	Auth_V1.GET("/growthstandards/:skip/:limit", ApiRoute.GetGrowthStandards)
	Auth_V1.POST("/person/add", ApiRoute.CreatePerson)
	Auth_V1.GET("/userinfo", ApiRoute.GetUserInfo)
	Auth_V1.GET("/persons", ApiRoute.GetPersonList)
	Auth_V1.GET("/news/:skip/:limit", ApiRoute.GetNews)
	Auth_V1.GET("/new/:mid", ApiRoute.GetNew)
	Auth_V1.POST("/updateperson", ApiRoute.SetPerson)

	//ignore auth
	NotAuth_V1 := V1.Group("/")
	NotAuth_V1.POST("/login", ApiRoute.Login)
	NotAuth_V1.POST("/wxlogin", ApiRoute.Wechat)

}
