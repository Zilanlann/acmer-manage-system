package global

import (
	"sync"

	"go.uber.org/zap"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	REDIS  redis.UniversalClient
	VP     *viper.Viper
	LOG    *zap.Logger
	lock   sync.RWMutex
)
