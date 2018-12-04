package Lib

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func LoadConfig(router *gin.Engine,confPath string){
	conf := &Yaml{}
	yamlFile, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Fatal(err)
	}else {
		err = yaml.Unmarshal(yamlFile, conf)
		if err!=nil{
			log.Fatal(err)
		}
	}
	router.Use(conf.ReadConfig)
}



func (conf *Yaml)ReadConfig(c *gin.Context)  {

	c.Set(conf.ConfigKey, conf)
	c.Next()
}

func GetConfig() (config *Yaml) {
	conf := &Yaml{}
	yamlFile, err := ioutil.ReadFile(CONFPATH)
	if err != nil {
		log.Fatal(err)
	}else {
		err = yaml.Unmarshal(yamlFile, conf)
		if err!=nil{
			log.Fatal(err)
		}
	}
	return conf
}



func SetLog()  {
	gin.DisableConsoleColor()
	//gin.ErrorLogger()

	// Logging to a file.
	f, _ := os.Create("server.log")
	gin.DefaultWriter = io.MultiWriter(f)

}