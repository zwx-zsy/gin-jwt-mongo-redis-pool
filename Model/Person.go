package Models

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

const collectionName = "SS_Person"

type Person struct {
	Name string `bson:"Name"`
	PassWord string `bson:"PassWord"`
}

type Men []Person


func Persons(c *gin.Context) *mgo.Collection {
	db, _ := c.Get("db")
	return db.(*mgo.Database).C(collectionName)
}



