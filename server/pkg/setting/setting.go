package setting

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

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

// Setup initialize the configuration instance
func Setup() {
	viper.SetConfigFile("config.toml")
	viper.AddConfigPath("/Users/hanyifan/Developer/acmer-manage-system/server/conf")
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
	ServerSetting.ReadTimeout = time.Second * ServerSetting.ReadTimeout
	ServerSetting.WriteTimeout = time.Second * ServerSetting.WriteTimeout
	return nil
}
