package Api

import (
	"TimeLine/Lib"
	M "TimeLine/Model"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"time"
)

// 登录参数
type PersonParam struct {
	NickName string `form:"nickname" json:"nickname" binding:"required"`
	Sex      int    `form:"sex" json:"sex" binding:"-"`
	Birthday string `form:"birthday" json:"birthday" binding:"required"`
	Born     int    `form:"born" json:"born" binding:"-"`
	Role     int    `form:"role" json:"role" binding:"-"`
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
	claims, _ := Lib.GetPayLoad(ctx)
	var PersonP PersonParam
	if err := ctx.ShouldBindJSON(&PersonP); err == nil {
		result := M.Person{Id: bson.NewObjectId(), NickName: PersonP.NickName,
			Sex: PersonP.Sex, Birthday: PersonP.Birthday, Born: PersonP.Born, Role: PersonP.Role, OpenId: claims.Payload.OpenId, CreateDateTime: time.Now()}
		err := M.Persons().Insert(&result)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": http.StatusPaymentRequired, "msg": err.Error()})
		} else {
			userone := &M.User{}
			user := M.User{WxOpenId: claims.Payload.OpenId, PersonId: result.Id.Hex()}
			err := M.Users().Find(bson.M{"WxOpenId": claims.Payload.OpenId}).One(&userone)
			if err != nil {
				user.CreateDateTime = time.Now()
			}
			ups, err := M.Users().Upsert(bson.M{"WxOpenId": claims.Payload.OpenId}, user)
			if err != nil {
				defer M.Rollback(M.CollectionName_Person, result.Id)
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success", "data": ups.UpsertedId})
			}
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": http.StatusPaymentRequired, "msg": err.Error()})
	}
}

//获取用户信息

type UserResult struct {
	User M.Person
	Days int
	List []M.Person
}

func GetUserInfo(c *gin.Context) {
	claims, _ := Lib.GetPayLoad(c)
	userone := &M.User{}
	result := &UserResult{}
	err := M.Users().Find(bson.M{"WxOpenId": claims.Payload.OpenId}).One(&userone)
	if err != nil {
		if err.Error() == "not found" {
			result.User = M.Person{}
			result.Days = -1
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": http.StatusText(http.StatusInternalServerError)})
		}
	} else {
		loc, _ := time.LoadLocation("Local")
		if userone.PersonId != "" {
			err := M.Persons().FindId(bson.ObjectIdHex(userone.PersonId)).One(&result.User)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": http.StatusText(http.StatusInternalServerError)})
			}
			toBeCharge := result.User.Birthday + " 00:00:00"
			parse_str_time, _ := time.ParseInLocation("2006-01-02 15:04:05", toBeCharge, loc)
			result.Days = Lib.TimeSub(time.Now(), parse_str_time)
		}
	}
	errs := M.Persons().Find(bson.M{"OpenId": claims.Payload.OpenId}).All(&result.List)
	if errs != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError,
			"msg": http.StatusText(http.StatusInternalServerError)})
	} else {
		if result.List == nil {
			result.List = []M.Person{}
		}
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success", "data": result, "Total": len(result.List)})
	}

}

func GetPersonList(c *gin.Context) {
	claims, _ := Lib.GetPayLoad(c)
	result := []M.Person{}
	err := M.Persons().Find(bson.M{"OpenId": claims.Payload.OpenId}).All(&result)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError,
			"msg": http.StatusText(http.StatusInternalServerError)})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success", "data": result, "Total": len(result)})
	}
}
