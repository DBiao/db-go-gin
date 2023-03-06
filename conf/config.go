package config

type Config struct {
	System System `mapstructure:"system" yaml:"system"`
	Zap    Zap    `mapstructure:"zap" yaml:"zap"`
	Mysql  Mysql  `mapstructure:"mysql" yaml:"mysql"`
	Redis  Redis  `mapstructure:"redis" yaml:"redis"`
}
