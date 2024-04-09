package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/zilanlann/acmer-manage-system/server/pkg/setting"
)

var ctx = context.Background()
var rdb *redis.Client

func Setup() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     setting.RedisSetting.Host,
		Password: setting.RedisSetting.Password,
		DB:       setting.RedisSetting.DB,
	})
}
