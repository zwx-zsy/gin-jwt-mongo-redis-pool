package Models

import (
	"TimeLine/Lib"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const CollectionName_User string = "bb_User"

type User struct {
	Id             bson.ObjectId `bson:"_id,omitempty"`
	WxOpenId       string        `bson:"WxOpenId"`
	UserCode       string        `bson:"UserCode"`
	PersonId       string        `bson:"PersonId"`
	CreateDateTime time.Time     `bson:"CreateDateTime"`
}

func Users() *mgo.Collection {
	//db, _ := c.Get("db")
	return Lib.DB.C(CollectionName_User)
}
