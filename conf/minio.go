package config

type Minio struct {
	Bucket   string `mapstructure:"bucket" yaml:"bucket"`       //
	Endpoint string `mapstructure:"endpoint" yaml:"endpoint"`   //
	AccessId string `mapstructure:"access-id" yaml:"access-id"` //
	Secret   string `mapstructure:"secret" yaml:"secret"`       // 密码
	UseHttps bool   `mapstructure:"use-https" yaml:"use-https"` //
}
