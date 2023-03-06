package config

type Redis struct {
	Host        string `mapstructure:"host" yaml:"host"`         // Ip地址
	Port        string `mapstructure:"port" yaml:"port"`         // 端口
	Password    string `mapstructure:"password" yaml:"password"` // 密码
	DB          int    `mapstructure:"db" yaml:"db"`
	PoolSize    int    `mapstructure:"pool-size" yaml:"pool-size"`       // 池子大小
	MinIdleConn int    `mapstructure:"minIdle-conn" yaml:"minIdle-conn"` // 最小连接
}
