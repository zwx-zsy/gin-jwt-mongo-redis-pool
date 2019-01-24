package Lib

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

var DB *mgo.Database

type Middleware struct {
	session *mgo.Session
}

func (this *MongoDB) String() string {
	return fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", this.User, this.Password, this.Host, this.Port, this.AuthDBName)
}

func NewMiddleware(session *mgo.Session) *Middleware {
	return &Middleware{
		session: session,
	}
}

func (m *Middleware) Connect(c *gin.Context) {
	s := m.session.Clone()
	DB = s.DB(ServerConf.DBConf.DatabaseName)
	defer s.Close()
	c.Next()
}

func Dial(router *gin.Engine) {
	Session, err := mgo.Dial(ServerConf.DBConf.String())
	if err != nil {
		log.Fatal(err)
	}
	Session.SetMode(mgo.Eventual, true)
	if err != nil {
		log.Fatal(err)
	}
	// middleware
	middleware := NewMiddleware(Session)
	router.Use(middleware.Connect)

}
