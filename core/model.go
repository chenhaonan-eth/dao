package core

type Config struct {
	Feishu Feishu `mapstructure:"feishu" json:"feishu" yaml:"feishu"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
}

type System struct {
	Port string `mapstructure:"port" json:"port" yaml:"port"`
	Dsn  string `mapstructure:"dsn" json:"dsn" yaml:"dsn"`
}
type Feishu struct {
	AppID             string `mapstructure:"app_id" json:"app_id" yaml:"app_id"`
	AppSecret         string `mapstructure:"app_secret" json:"app_secret" yaml:"app_secret"`
	EncryptKey        string `mapstructure:"encrypt_key" json:"encrypt_key" yaml:"encrypt_key"`
	VerificationToken string `mapstructure:"verification_token" json:"verification_token" yaml:"verification_token"`
}
type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                           // 级别
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                        // 输出
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                        // 日志前缀
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`                 // 日志文件夹
	ShowLine      bool   `mapstructure:"show-line" json:"showLine" yaml:"showLine"`                 // 显示行
	EncodeLevel   string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" yaml:"stacktrace-key"` // 栈名
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`
}
