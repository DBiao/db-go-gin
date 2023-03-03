package config

type System struct {
	Port      int    `mapstructure:"port" yaml:"port"`             // 端口值
	Mode      string `mapstructure:"mode" yaml:"mode"`             // gin调式模式
	ModelPath string `mapstructure:"model-path" yaml:"model-path"` // 存放casbin模型的相对路径
}
