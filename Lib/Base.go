package Lib

const (
	CONFPATH string = "/septnet/config/conf.yaml"
	CONFKEY string = "Config"
)

type Yaml struct {
	ConfigKey string `yaml:"ConfigKey"`
	DBConf MongoDB `yaml:"mongodb"`
	RedisConf Redis `yaml:"redis"`
	RedisConn string `yaml:"redisConn"`
}


type MongoDB struct {
	User string `yaml:"db_user"`
	Host string `yaml:"db_host"`
	Password string `yaml:"db_pass"`
	Port string `yaml:"db_port"`
	DatabaseName string `yaml:"db_database_name"`
	Uri string `yaml:"url"`
}