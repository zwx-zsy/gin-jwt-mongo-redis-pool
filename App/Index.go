package Api

import (
	"TimeLine/Lib"
	M "TimeLine/Model"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// 登录参数
type PersonParam struct {
	NickName string `form: "nickname" json: "nickname" binding: "required"`
	Sex      int    `form: "sex" json: "sex" binding: "-"`
	Birthday string `form: "birthday" json: "birthday" binding: "required"`
	Born     int    `form: "born" json: "born" binding: "-"`
	Role     int    `form: "role" json: "role" binding: "-"`
}

func GetGrowthStandards(c *gin.Context) {
	// 获取成长标准说明
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
type Users M.User

func (t Users) String() string {
	return fmt.Sprintf("学号: %s\n真实姓名: %s\n年龄: %s\n", t.PersonId, t.WxOpenId, t.Id)
}
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
				result.Days = Lib.TimeSub(
					time.Now(), parse_str_time)
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
