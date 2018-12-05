package Lib

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var ServerConf *Yaml

func LoadConfig(router *gin.Engine, confPath string) {

	yamlFile, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Fatal(err)
	} else {
		err = yaml.Unmarshal(yamlFile, &ServerConf)
		if err != nil {
			log.Fatal(err)
		}
	}
	//router.Use(ServerConf.ReadConfig)
}

//func (conf *ServerConf)ReadConfig(c *gin.Context)  {
//
//	//c.Set(conf.ConfigKey, conf)
//	c.Next()
//}

//func GetConfig() (config *Yaml) {
//	//conf := &Yaml{}
//	yamlFile, err := ioutil.ReadFile(CONFPATH)
//	if err != nil {
//		log.Fatal(err)
//	}else {
//		err = yaml.Unmarshal(yamlFile, ServerConf)
//		if err!=nil{
//			log.Fatal(err)
//		}
//	}
//	return ServerConf
//}

func SetLog() {
	gin.DisableConsoleColor()
	//gin.ErrorLogger()

	// Logging to a file.
	f, _ := os.Create("server.log")
	gin.DefaultWriter = io.MultiWriter(f)

}
