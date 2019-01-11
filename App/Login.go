package Api

import (
	"TimeLine/Lib"
	M "TimeLine/Model"
	"encoding/json"
	Jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"time"
)

// 登录参数
type LoginParam struct {
	User     string `form:"user" json:"user" binding:"-"`
	Password string `form:"password" json:"password" binding:"required"`
}

//微信 code登录

type WeixinCode struct {
	Code string `form:"code" json:"code" binding:"required"`
}

//登录操作
func Login(c *gin.Context) {
	var jsons LoginParam
	if err := c.ShouldBindJSON(&jsons); err == nil {
		result := M.Person{}
		err := M.Persons().Find(bson.M{"Name": jsons.User}).One(&result)
		if err != nil {
			switch err {
			case mgo.ErrNotFound:
				c.AbortWithStatusJSON(http.StatusOK, gin.H{
					"code": http.StatusNotFound,
					"msg":  err.Error(),
				})
			default:

				c.JSON(http.StatusOK, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  err.Error(),
				})
			}
		} else {
			if jsons.User == result.NickName && jsons.Password == result.Birthday {
				generateToken(c, Lib.Payload{Name: jsons.User})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code": http.StatusBadRequest,
					"msg":  "用户名或密码错误！",
				})
			}
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusPaymentRequired, "msg": err.Error()})
	}
}

func generateToken(c *gin.Context, payload Lib.Payload) string {
	j := &Lib.JWT{
		[]byte(Lib.SignKey),
	}
	claims := Lib.CustomClaims{
		Payload: payload,
		StandardClaims: Jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() + Lib.ServerConf.JwtConf.Notbefore), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + Lib.ServerConf.JwtConf.Exptime),   // 过期时间 一小时*24
			Issuer:    Lib.ServerConf.JwtConf.Issuer,                               //签名的发行者
		},
	}

	if token, err := j.CreateToken(claims); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return ""
	} else {
		//c.JSON(http.StatusOK, gin.H{
		//	"code": http.StatusOK,
		//	"msg":  http.StatusText(http.StatusOK),
		//	"data": token})
		return token
	}

}

type wxOpenid struct {
	Sessionkey string `json:"session_key"`
	Openid     string `json:"openid"`
	Errcode    int32  `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

//code换 openid
func (wxc *WeixinCode) codetoopenid() (res *wxOpenid, e error) {
	client := &http.Client{}

	//get url
	wechat := Lib.ServerConf.WeChat
	url := wechat.CodeUrl(wxc.Code)
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	//处理返回结果
	response, _ := client.Do(reqest)
	//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
	//body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	resp := &wxOpenid{}
	err = json.Unmarshal(body, resp)
	return &wxOpenid{
		resp.Sessionkey,
		resp.Openid,
		resp.Errcode,
		resp.Errmsg}, err
}

func Wechat(ctx *gin.Context) {
	/**
	微信方式登录
	**/
	weixincode := WeixinCode{}
	if err := ctx.ShouldBindJSON(&weixincode); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "msg": http.StatusText(http.StatusBadRequest)})
	} else {
		body, err := weixincode.codetoopenid()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{"code": http.StatusInternalServerError, "msg": err.Error()})
		} else {
			if body.Errcode > 0 {
				ctx.AbortWithStatusJSON(http.StatusOK,
					gin.H{"code": body.Errcode, "msg": body.Errmsg})
			} else {
				userone := &M.User{}
				err := M.Users().Find(bson.M{"WxOpenId": body.Openid}).One(&userone)
				if err != nil {
					if userone.WxOpenId == "" {
						userone.WxOpenId = body.Openid
						userone.Id = bson.NewObjectId()
						userone.CreateDateTime = time.Now()
						M.Users().Insert(userone)
					}
				}
				token := generateToken(ctx, Lib.Payload{OpenId: body.Openid})
				ctx.JSON(http.StatusOK, gin.H{
					"code":      http.StatusOK,
					"msg":       http.StatusText(http.StatusOK),
					"data":      token,
					"personmid": userone.PersonId})
			}
		}
	}
}
