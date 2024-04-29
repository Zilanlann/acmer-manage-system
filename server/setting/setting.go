package setting

import (
	"fmt"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

type Cfg struct {
	Server   Server   `mapstructure:"server"`
	Jwt      Jwt      `mapstructure:"jwt"`
	Database Database `mapstructure:"database"`
	Redis    Redis    `mapstructure:"redis"`
	Tencent  Tencent  `mapstructure:"tencent"`
	Zap      Zap      `mapstructure:"zap"`
}

var CfgSetting = &Cfg{}

type Server struct {
	RunMode      string        `mapstructure:"run_mode"`
	HttpPort     int           `mapstructure:"http_port"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

var ServerSetting = &Server{}

type Database struct {
	Type         string `mapstructure:"type"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	Name         string `mapstructure:"name"`
	TablePrefix  string `mapstructure:"table_prefix"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}

var DatabaseSetting = &Database{}

type Jwt struct {
	Secret           string        `mapstructure:"secret"`
	LongExpiresTime  time.Duration `mapstructure:"long_expires_time"`
	ShortExpiresTime time.Duration `mapstructure:"short_expires_time"`
	Issuer           string        `mapstructure:"issuer"`
}

var JwtSetting = &Jwt{}

type Redis struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

var RedisSetting = &Redis{}

type Tencent struct {
	SecretId   string `mapstructure:"secret_id"`
	SecretKey  string `mapstructure:"secret_key"`
	MailTempID uint64 `mapstructure:"mail_temp_id"`
	MailFrom   string `mapstructure:"mail_from"`
}

var TencentSetting = &Tencent{}

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                            // 级别
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 日志前缀
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                         // 输出
	Dir           string `mapstructure:"dir" json:"dir"  yaml:"dir"`                                 // 日志文件夹
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"` // 栈名
	MaxAge        int    `mapstructure:"max-age" json:"max-age" yaml:"max-age"`                      // 日志留存时间
	ShowLine      bool   `mapstructure:"show-line" json:"show-line" yaml:"show-line"`                // 显示行
	LogInConsole  bool   `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"` // 输出控制台
}

var ZapSetting = &Zap{}

// ZapEncodeLevel 根据 EncodeLevel 返回 zapcore.LevelEncoder
func (z *Zap) ZapEncodeLevel() zapcore.LevelEncoder {
	switch {
	case z.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		return zapcore.LowercaseLevelEncoder
	case z.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case z.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		return zapcore.CapitalLevelEncoder
	case z.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

// TransportLevel 根据字符串转化为 zapcore.Level
func (z *Zap) TransportLevel() zapcore.Level {
	z.Level = strings.ToLower(z.Level)
	switch z.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

// Setup initialize the configuration instance
func Setup() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("conf")
	viper.SetEnvPrefix("AMS")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	if err := unmarshal(); err != nil {
		panic(fmt.Errorf("fatal error bind config: %w", err))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := unmarshal(); err != nil {
			panic(fmt.Errorf("fatal error bind config: %w", err))
		}
	})
}

// unmarshal config to structs
func unmarshal() error {
	// 由于viper仍未解决的解组bug,只能用Unmarshal而不是UnmarshalKey来解组才能支持自动读取环境变量
	if err := viper.Unmarshal(&CfgSetting); err != nil {
		return err
	}
	ServerSetting = &CfgSetting.Server
	JwtSetting = &CfgSetting.Jwt
	DatabaseSetting = &CfgSetting.Database
	RedisSetting = &CfgSetting.Redis
	TencentSetting = &CfgSetting.Tencent
	ZapSetting = &CfgSetting.Zap
	ServerSetting.ReadTimeout = time.Second * ServerSetting.ReadTimeout
	ServerSetting.WriteTimeout = time.Second * ServerSetting.WriteTimeout
	JwtSetting.LongExpiresTime = time.Hour * JwtSetting.LongExpiresTime
	JwtSetting.ShortExpiresTime = time.Hour * JwtSetting.ShortExpiresTime
	return nil
}
