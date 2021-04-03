package config

// python goods-srv  配置信息
type GoodsSrvConfig struct {
	Host       string `mapstructure:"host" json:"host"`
	Port       int    `mapstructure:"port" json:"port"`
	ServerName string `mapstructure:"server_name" json:"server_name"`
}

type ServerConfig struct {
	Name         string         `mapstructure:"name" json:"name"`
	Port         int            `mapstructure:"port" json:"port"`
	Host         string         `mapstructure:"host" json:"host"`
	Tags         []string       `mapstructure:"tags" json:"tags"`
	GoodsSrvInfo GoodsSrvConfig `mapstructure:"goods_srv" json:"goods_srv"`
	JWTInfo      JWTConfig      `mapstructure:"jwt" json:"jwt"`
	RedisInfo    RedisConfig    `mapstructure:"redis" json:"redis"`
	ConsulInfo   ConsulConfig   `mapstructure:"consul" json:"consul"`
}

type RedisConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
	Expire     int    `mapstructure:"expire" json:"expire"`
	Issuer     string `mapstructure:"issuer" json:"issuer"`
}
type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host" json:"host"`
	Port      int    `mapstructure:"port" json:"port"`
	NameSpace string `mapstructure:"namespace" json:"namespace"`
	User      string `mapstructure:"user" json:"user"`
	Password  string `mapstructure:"password" json:"password"`
	DataId    string `mapstructure:"dataid" json:"dataid"`
	Group     string `mapstructure:"group" json:"group"`
}
