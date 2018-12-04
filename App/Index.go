package Api

import (
	"TimeLine/Lib"
	M "TimeLine/Model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloServer(c *gin.Context) {
	//if _, err := Lib.Set(c, "testkey", "1234567890");err!=nil{
	//	c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"code":http.StatusInternalServerError,
	//		"msg":http.StatusText(http.StatusInternalServerError)})
	//
	//}
	//if replyS, err := Lib.Get(c, "testkey");err!=nil{
	//	c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"code":http.StatusInternalServerError,
	//		"msg":http.StatusText(http.StatusInternalServerError)})
	//}else {
	//	fmt.Println(replyS)
	//	//获取到 redis中的指定的 key的 value值 做对应的操作
	//}
	claims, b := Lib.GetPayLoad(c)
	payload := Lib.CustomClaims{}.Payload
	if b {
		payload = claims.Payload
	}
	document := M.Person{Name: payload.Name, PassWord: payload.Name}
	errs := M.Persons(c).Insert(&document)
	if errs != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"code":http.StatusInternalServerError,
			"msg":http.StatusText(http.StatusInternalServerError)})
	}else{
		c.JSON(200, gin.H{
			"code":    200,
			"msg": "success",
			"data":    fmt.Sprintf("Hello %s !",payload.Name),
		})
	}
}
