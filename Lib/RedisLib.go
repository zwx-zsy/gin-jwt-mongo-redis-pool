package Lib

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"time"
)


type RMiddleware struct {
	pool *redis.Pool
}

type Redis struct {
	Host string `yaml:"redis_host"`
	Password string `yaml:"redis_pass"`
	Port string `yaml:"redis_port"`
	DatabaseName string `yaml:"redis_db"`
}


func (this *Redis) String() string {
	return fmt.Sprintf("%s:%s", this.Host, this.Port)
}

func RedisPool(router *gin.Engine) {
	config := GetConfig().RedisConf
	newPool := NewPool(config.String(),config)

	router.Use(newPool.RedisConn)
}

func NewPool(addr string,config Redis) *RMiddleware {
	pool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, e := redis.Dial("tcp", addr)
			if e != nil {
				return nil, e
			}
			//设置密码
			if _, e = conn.Do("AUTH", config.Password); e != nil {
				conn.Close()
				return nil, e
			}
			//选择 db
			if _, err := conn.Do("SELECT", config.DatabaseName); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, nil
		},
	}
	return &RMiddleware{pool:pool}
}

func (p *RMiddleware)RedisConn(c *gin.Context)  {
	connS := p.pool.Get()
	defer connS.Close()
	config, _ := c.Get(CONFKEY)
	c.Set(config.(*Yaml).RedisConn,connS)
	c.Next()
}

func getRedis(c *gin.Context) (redis.Conn,bool) {
	config, _ := c.Get(CONFKEY)
	if value, exists := c.Get(config.(*Yaml).RedisConn);exists{
		return value.(redis.Conn),true
	}else{
		return nil,false
	}
}

func Set(c *gin.Context, args ...interface{}) (reply interface{}, err error) {
	if conn, b := getRedis(c);b{
		reply, err := conn.Do("SET", args...)
		return reply,err
	}else{
		return nil,errors.New("redisConn not found")
	}
}

func Get(c *gin.Context,args ...interface{})(reply interface{},err error)  {
	conn, b := getRedis(c)
	if b == false{
		return nil,errors.New("redisConn not found")
	}else{
		if s, err := redis.String(conn.Do("GET", args...));err!=nil{
			return nil,err

		}else{
			return s,nil
		}
	}

}