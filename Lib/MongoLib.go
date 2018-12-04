package Lib

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"log"
)

type Middleware struct {
	session *mgo.Session
}


func (this *MongoDB) String() string {
	return fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", this.User, this.Password, this.Host, this.Port, this.DatabaseName)
}

func NewMiddleware(session *mgo.Session) *Middleware {
	return &Middleware{
		session: session,
	}
}

func (m *Middleware)Connect(c *gin.Context) {
	s := m.session.Clone()
	db := s.DB("test")
	defer s.Close()
	c.Set("db", db)
	c.Next()
}

func Dial(router *gin.Engine){
	config := GetConfig().DBConf

	//valueOf := reflect.ValueOf(config)
	//for i := 0; i < valueOf.NumField(); i++ {
	//	fmt.Println(i,valueOf.Field(i))
	//}

	Session ,err := mgo.Dial(config.String())
	if err !=nil{
		log.Fatal(err)
	}
	Session.SetMode(mgo.Eventual,true)
	if err != nil {
		log.Fatal(err)
	}
	// middleware
	middleware := NewMiddleware(Session)
	router.Use(middleware.Connect)

}
