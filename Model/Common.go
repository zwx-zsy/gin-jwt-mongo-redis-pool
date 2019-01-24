package Models

import (
	"TimeLine/Lib"

	"gopkg.in/mgo.v2/bson"
)

//type CustomCollection mgo.Collection
//
//type CustomContext gin.Context

//func (m *CustomCollection)() *mgo.Collection {
//	db, _ := m.Get("db")
//	dbfmt := db.(*mgo.Database)
//	m.s()
//}

func Rollback(collectionName string, id bson.ObjectId) (removeid bson.ObjectId,err error){
	//由于没有事务处理这个做一个撤销操作

	errs := Lib.DB.C(collectionName).RemoveId(id)
	//log.Fatalf("%v",e)
	if errs != nil{
		return "",errs
	}else {
		return id,nil
	}
}
