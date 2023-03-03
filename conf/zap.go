package config

type Zap struct {
	Level         string `mapstructure:"level" yaml:"level"`                   // 级别
	Format        string `mapstructure:"format" yaml:"format"`                 // 输出
	Prefix        string `mapstructure:"prefix" yaml:"prefix"`                 // 日志前缀
	ShowLine      bool   `mapstructure:"show-line" yaml:"showLine"`            // 显示行
	EncodeLevel   string `mapstructure:"encode-level" yaml:"encode-level"`     // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" yaml:"stacktrace-key"` // 栈名
	LinkName      string `mapstructure:"link-name" yaml:"link-name"`           // 软链接名称
	LogInConsole  bool   `mapstructure:"log-in-console" yaml:"log-in-console"` // 输出控制台
}
