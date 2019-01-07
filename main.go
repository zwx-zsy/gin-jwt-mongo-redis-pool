package main

import (
	"TimeLine/Lib"
	"TimeLine/Router"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	//gin.SetMode(gin.DebugMode)
	// 初始化服务
	router := gin.New()
	//获取配置文件配置信息

	Lib.LoadConfig(router, Lib.CONFPATH)
	//实例化数据库操作

	Lib.Dial(router)

	//redis
	Lib.RedisPool(router)
	Lib.SetLog()

	// Registered routing
	Router.RegisterRouter(router)
	router.Run(fmt.Sprintf("%s:%s", Lib.ServerConf.Server.Host, Lib.ServerConf.Server.Port))
}
