package redis

import (
	"context"
	"crypto/tls"

	"github.com/redis/go-redis/v9"
	"github.com/zilanlann/acmer-manage-system/server/pkg/setting"
)

var Ctx = context.Background()
var RDB *redis.Client

func Setup() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     setting.RedisSetting.Host,
		Password: setting.RedisSetting.Password,
		DB:       setting.RedisSetting.DB,
		TLSConfig: &tls.Config{},
	})
}
