package Api

import (
	M "TimeLine/Model"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

// 登录参数
type PersonParam struct {
	NickName string `form:"nickname" json:"nickname" binding:"required"`
	Sex      int    `form:"sex" json:"sex" binding:"-"`
	Birthday string `form:"birthday" json:"birthday" binding:"required"`
}

func GetGrowthStandards(c *gin.Context) {
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
	//claims, b := Lib.GetPayLoad(c)
	//payload := Lib.CustomClaims{}.Payload
	//if b {
	//	payload = claims.Payload
	//}
	gs := []M.GrowthStandard{}
	e := M.GrowthStandards().Find(bson.M{"Type": 0}).Skip(10).Limit(20).All(&gs)

	if e != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError,
			"msg": http.StatusText(http.StatusInternalServerError)})
	} else {
		c.JSON(200, gin.H{"code": 200, "msg": "success", "data": gs, "Total": len(gs)})
	}

}

func CreatePerson(ctx *gin.Context) {
	var PersonP PersonParam
	if err := ctx.ShouldBindJSON(&PersonP); err == nil {
		result := M.Person{Id: bson.NewObjectId(), NickName: PersonP.NickName, Sex: PersonP.Sex, Birthday: PersonP.Birthday}
		err := M.Persons().Insert(&result)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": 402, "msg": err.Error()})
		} else {
			ctx.JSON(200, gin.H{"code": 200, "msg": "success", "data": result.Id})
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": 402, "msg": err.Error()})
	}
}
