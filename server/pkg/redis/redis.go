package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/zilanlann/acmer-manage-system/server/setting"
)

var Ctx = context.Background()
var RDB *redis.Client

func Setup() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     setting.RedisSetting.Host,
		Password: setting.RedisSetting.Password,
		DB:       setting.RedisSetting.DB,
	})
}
