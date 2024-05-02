package global

import (
	"sync"

	"go.uber.org/zap"

	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	REDIS  *redis.Client
	VP     *viper.Viper
	LOG    *zap.Logger
	Casbin *casbin.Enforcer
	lock   sync.RWMutex
)
