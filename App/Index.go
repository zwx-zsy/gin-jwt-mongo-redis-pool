package Api

import (
	M "TimeLine/Model"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	strconv "strconv"
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
	skip, _ := strconv.Atoi(c.Param("skip"))
	limit, _ := strconv.Atoi(c.Param("limit"))
	gs := []M.GrowthStandard{}
	e := M.GrowthStandards().Find(nil).Sort("Type", "Days").Skip(skip).Limit(limit).All(&gs)

	if e != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError,
			"msg": http.StatusText(http.StatusInternalServerError)})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success", "data": gs, "Total": len(gs)})
	}

}

func CreatePerson(ctx *gin.Context) {
	var PersonP PersonParam
	if err := ctx.ShouldBindJSON(&PersonP); err == nil {
		result := M.Person{Id: bson.NewObjectId(), NickName: PersonP.NickName, Sex: PersonP.Sex, Birthday: PersonP.Birthday}
		err := M.Persons().Insert(&result)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": http.StatusPaymentRequired, "msg": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success", "data": result.Id})
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": http.StatusPaymentRequired, "msg": err.Error()})
	}
}
