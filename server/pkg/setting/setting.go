package setting

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type         string
	User         string
	Password     string
	Host         string
	Name         string
	TablePrefix  string
	MaxIdleConns int
	MaxOpenConns int
}

var DatabaseSetting = &Database{}

type Jwt struct {
	Secret           string
	LongExpiresTime  time.Duration
	ShortExpiresTime time.Duration
	Issuer           string
}

var JwtSetting = &Jwt{}

type Redis struct {
	Host     string
	Password string
	DB       int
}

var RedisSetting = &Redis{}

type Tencent struct {
	SecretId  string
	SecretKey string
	MailTempID  uint64
	MailFrom  string
}

var TencentSetting = &Tencent{}

// Setup initialize the configuration instance
func Setup() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("conf")
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
	if err := viper.UnmarshalKey("Server", &ServerSetting); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("Database", &DatabaseSetting); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("Jwt", &JwtSetting); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("Redis", &RedisSetting); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("Tencent", &TencentSetting); err != nil {
		return err
	}
	ServerSetting.ReadTimeout = time.Second * ServerSetting.ReadTimeout
	ServerSetting.WriteTimeout = time.Second * ServerSetting.WriteTimeout
	JwtSetting.LongExpiresTime = time.Hour * JwtSetting.LongExpiresTime
	JwtSetting.ShortExpiresTime = time.Hour * JwtSetting.ShortExpiresTime
	return nil
}
