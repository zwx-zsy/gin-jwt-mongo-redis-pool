package Lib

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	yaml "gopkg.in/yaml.v2"
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

func SetLog() {
	gin.DisableConsoleColor()
	gin.ErrorLogger()

	// Logging to a file.
	f, _ := os.Create("server.log")
	gin.DefaultWriter = io.MultiWriter(f)

}
