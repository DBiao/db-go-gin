package config

type Mysql struct {
	Path        string `mapstructure:"path" yaml:"path"`                   // 服务器地址:端口
	Config      string `mapstructure:"config" yaml:"config"`               // 高级配置
	Dbname      string `mapstructure:"db-name" yaml:"db-name"`             // 数据库名
	Username    string `mapstructure:"username" yaml:"username"`           // 数据库用户名
	Password    string `mapstructure:"password" yaml:"password"`           // 数据库密码
	MaxIdleConn int    `mapstructure:"max-idle-conn" yaml:"max-idle-conn"` // 空闲中的最大连接数
	MaxOpenConn int    `mapstructure:"max-open-conn" yaml:"max-open-conn"` // 打开到数据库的最大连接数
	LogMode     string `mapstructure:"log-mode" yaml:"log-mode"`           // 是否开启Gorm全局日志
	LogZap      bool   `mapstructure:"log-zap" yaml:"log-zap"`             // 是否通过zap写入日志文件
}
