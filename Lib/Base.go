package Lib

import (
	"fmt"
	"time"
)

const (
	CONFPATH string = "/etc/tl/conf.yaml" // 配置文件地址
	// CONFPATH string = "/septnet/config/conf.yaml" // 配置文件地址
	CONFKEY string = "Config" //配置文件的key
)

//yaml的结构

//start
type Yaml struct {
	ConfigKey string  `yaml:"ConfigKey"`
	DBConf    MongoDB `yaml:"mongodb"`
	RedisConf Redis   `yaml:"redis"`
	RedisConn string  `yaml:"redisConn"`
	JwtConf   JwtConf `yaml:"JwtConf"`
	WeChat    WeChats `yaml:"wechat"`
	Server    Server  `yaml:"Server"`
}

type Server struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
}
type MongoDB struct {
	User         string `yaml:"db_user"`
	Host         string `yaml:"db_host"`
	Password     string `yaml:"db_pass"`
	Port         string `yaml:"db_port"`
	DatabaseName string `yaml:"db_database_name"`
	AuthDBName   string `yaml:"db_auth_name"`
	Uri          string `yaml:"url"`
}

type JwtConf struct {
	Issuer    string `yaml:"issuer"`
	Exptime   int64  `yaml:"exptime"`
	Notbefore int64  `yaml:"notbefore"`
}

type WeChats struct {
	APPID     string `yaml:"APPID"`
	APPSECRET string `yaml:"APPSECRET"`
}

//end

func (this *WeChats) CodeUrl(code string) string {
	return fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v&grant_type=authorization_code", this.APPID, this.APPSECRET, code)
}

//计算二个时间的相差多少天
func TimeSub(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)
	return int(t1.Sub(t2).Hours() / 24)
}
