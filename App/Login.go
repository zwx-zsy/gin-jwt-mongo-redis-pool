package Api

import (
	"TimeLine/Lib"
	M "TimeLine/Model"
	Jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"net/http"
	"time"
)

// 登录参数
type LoginParam struct {
	User     string `form:"user" json:"user" binding:"-"`
	Password string `form:"password" json:"password" binding:"required"`
}


//登录操作
func Login(c *gin.Context) {
	var jsons LoginParam
	if err := c.ShouldBindJSON(&jsons); err == nil {
		result := M.Person{}
		err := M.Persons(c).Find(bson.M{"Name": jsons.User}).One(&result)
		if err != nil{
			switch err {
			case mgo.ErrNotFound:
				c.AbortWithStatusJSON(http.StatusOK, gin.H{
					"code": http.StatusNotFound,
					"msg": err.Error(),
				})
			default:

				c.JSON(http.StatusOK, gin.H{
					"code": http.StatusInternalServerError,
					"msg": err.Error(),
				})
			}
		}else {
			if jsons.User ==  result.Name&& jsons.Password == result.PassWord {
				generateToken(c, jsons.User)
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code": 400,
					"msg":    "用户名或密码错误！",
				})
			}
		}


	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code":402,"msg":err.Error()})
	}
}


func generateToken(c *gin.Context, user string) {
	j := &Lib.JWT{
		[]byte("www.vcoding.com"),
	}

	claims := Lib.CustomClaims{
		Payload:Lib.Payload{
			user,
			true},
		StandardClaims:Jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "newtrekWang",                   //签名的发行者
			},
	}

	if token, err := j.CreateToken(claims); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":    err.Error(),
		})
		return
	} else {

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "登录成功！",
			"data": token})
		return

	}

}
