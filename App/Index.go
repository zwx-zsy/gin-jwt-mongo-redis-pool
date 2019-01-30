package Api

import (
	"TimeLine/Lib"
	M "TimeLine/Model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
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
	// 获取成长标准说明
	skip, _ := strconv.Atoi(c.Param("skip"))
	limit, _ := strconv.Atoi(c.Param("limit"))
	born, _ := strconv.Atoi(c.Param("born"))
	gs := []M.GrowthStandard{}
	e := M.GrowthStandards().Find(bson.M{"Type": born}).Sort("Days").Skip(skip).Limit(limit).All(&gs)

	if e != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError,
			"msg": http.StatusText(http.StatusInternalServerError)})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success", "data": gs, "Total": len(gs)})
	}
}

//创建宝宝
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
type Users M.User

func (t Users) String() string {
	return fmt.Sprintf("学号: %s\n真实姓名: %s\n年龄: %s\n", t.PersonId, t.WxOpenId, t.Id)
}

//获取当前用户的信息包括绑定的宝宝
func GetUserInfo(c *gin.Context) {
	claims, err := Lib.GetPayLoad(c)
	if !err {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": http.StatusText(http.StatusInternalServerError)})
	} else {
		userone := &Users{PersonId: "sadasdfa", Id: "hsjkdfs"}
		result := &UserResult{}
		fmt.Println(userone)
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
				Days := Lib.TimeSub(
					time.Now(), parse_str_time)
				if Days < 0 {
					Days = -Days
				}
				result.Days = Days
			} else {
				result.User = M.Person{}
				result.Days = -1

			}
		}
		fmt.Println(userone)
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

}

//获取宝宝列表
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

//获取资讯列表
func GetNews(c *gin.Context) {
	skip, _ := strconv.Atoi(c.Param("skip"))
	limit, _ := strconv.Atoi(c.Param("limit"))
	result := []M.Message{}
	count, _ := M.Messages().Count()
	err := M.Messages().Find(bson.M{}).Skip(skip).Limit(limit).Select(bson.M{"id": 1, "Title": 1, "CreateDateTime": 1}).All(&result)

	r := rand.New(rand.NewSource(100000))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError,
			"msg": http.StatusText(http.StatusInternalServerError)})
	} else {
		res := make([]M.Message, 0)
		randlist := r.Perm(len(result))[:3]
		for index := range randlist {
			res = append(res, result[randlist[index]])
		}
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success", "data": res, "Total": count})
	}
}

//获取单个资讯内容
func GetNew(c *gin.Context) {
	Mid := c.Param("mid")
	result := &M.Message{}
	err := M.Messages().FindId(bson.ObjectIdHex(Mid)).One(&result)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError,
			"msg": http.StatusText(http.StatusInternalServerError)})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success", "data": result})
	}
}

//请求参数
type PersonIdP struct {
	PersonId string `json:"person_id" binding:"required"`
}

//修改绑定宝宝
func SetPerson(c *gin.Context) {
	claims, _ := Lib.GetPayLoad(c)
	var PersonP PersonIdP
	if e := c.ShouldBindJSON(&PersonP); e == nil {
		e := M.Users().Update(bson.M{"WxOpenId": claims.Payload.OpenId}, bson.M{"$set": bson.M{"PersonId": PersonP.PersonId}})
		if e != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError,
				"msg": http.StatusText(http.StatusInternalServerError)})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success", "data": PersonP.PersonId})
		}
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": http.StatusPaymentRequired, "msg": e.Error()})
	}

}
