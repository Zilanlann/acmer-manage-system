package setting

import (
	"fmt"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Cfg struct {
	Server   Server   `mapstructure:"server"`
	Jwt      Jwt      `mapstructure:"jwt"`
	Database Database `mapstructure:"database"`
	Redis    Redis    `mapstructure:"redis"`
	Tencent  Tencent  `mapstructure:"tencent"`
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
	// 由于viper解组bug,只能用Unmarshal来解组才能支持自动读取环境变量
	if err := viper.Unmarshal(&CfgSetting); err != nil {
		return err
	}
	ServerSetting = &CfgSetting.Server
	JwtSetting = &CfgSetting.Jwt
	DatabaseSetting = &CfgSetting.Database
	RedisSetting = &CfgSetting.Redis
	TencentSetting = &CfgSetting.Tencent
	ServerSetting.ReadTimeout = time.Second * ServerSetting.ReadTimeout
	ServerSetting.WriteTimeout = time.Second * ServerSetting.WriteTimeout
	JwtSetting.LongExpiresTime = time.Hour * JwtSetting.LongExpiresTime
	JwtSetting.ShortExpiresTime = time.Hour * JwtSetting.ShortExpiresTime
	return nil
}
