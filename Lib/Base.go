package Lib

const (
	CONFPATH string = "/septnet/config/conf.yaml" // 配置文件地址
	CONFKEY  string = "Config"                    //配置文件的key
)

//yaml的结构

//start
type Yaml struct {
	ConfigKey string  `yaml:"ConfigKey"`
	DBConf    MongoDB `yaml:"mongodb"`
	RedisConf Redis   `yaml:"redis"`
	RedisConn string  `yaml:"redisConn"`
	JwtConf   JwtConf `yaml:"JwtConf"`
}

type MongoDB struct {
	User         string `yaml:"db_user"`
	Host         string `yaml:"db_host"`
	Password     string `yaml:"db_pass"`
	Port         string `yaml:"db_port"`
	DatabaseName string `yaml:"db_database_name"`
	Uri          string `yaml:"url"`
}

type JwtConf struct {
	Issuer    string `yaml:"issuer"`
	Exptime   int64  `yaml:"exptime"`
	Notbefore int64  `yaml:"notbefore"`
}

//end
