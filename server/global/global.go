package global

import (
	"sync"

	"github.com/songzhibin97/gkit/cache/local_cache"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/zilanlann/acmer-manage-system/server/config"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	AMS_DB     *gorm.DB
	AMS_DBList map[string]*gorm.DB
	AMS_REDIS  *redis.Client
	AMS_CONFIG config.Server
	AMS_VP     *viper.Viper
	// AMS_LOG    *oplogging.Logger
	AMS_LOG                 *zap.Logger
	AMS_Concurrency_Control = &singleflight.Group{}

	BlackCache local_cache.Cache
	lock       sync.RWMutex
)
