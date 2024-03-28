package system

import (
	"context"
	"errors"
	"fmt"

	"github.com/gookit/color"
	"github.com/zilanlann/acmer-manage-system/server/config"

	"github.com/zilanlann/acmer-manage-system/server/utils"

	"github.com/gofrs/uuid/v5"
	"github.com/zilanlann/acmer-manage-system/server/global"
	"github.com/zilanlann/acmer-manage-system/server/model/system/request"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLInitHandler struct{}

func NewMysqlInitHandler() *MySQLInitHandler {
	return &MySQLInitHandler{}
}

// WriteConfig mysql回写配置
func (h MySQLInitHandler) WriteConfig(ctx context.Context) error {
	c, ok := ctx.Value("config").(config.MySQL)
	if !ok {
		return errors.New("mysql config invalid")
	}
	global.AMS_CONFIG.MySQL = c
	global.AMS_CONFIG.JWT.SigningKey = uuid.Must(uuid.NewV4()).String()
	cs := utils.StructToMap(global.AMS_CONFIG)
	for k, v := range cs {
		global.AMS_VP.Set(k, v)
	}
	return global.AMS_VP.WriteConfig()
}

// EnsureDB 创建数据库并初始化 mysql
func (h MySQLInitHandler) EnsureDB(ctx context.Context, conf *request.InitDB) (next context.Context, err error) {
	if s, ok := ctx.Value("dbtype").(string); !ok || s != "mysql" {
		return ctx, ErrDBTypeMismatch
	}

	c := conf.ToMysqlConfig()
	next = context.WithValue(ctx, "config", c)
	if c.Dbname == "" {
		return ctx, nil
	} // 如果没有数据库名, 则跳出初始化数据

	dsn := conf.MysqlEmptyDsn()
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", c.Dbname)
	if err = createDatabase(dsn, "mysql", createSql); err != nil {
		return nil, err
	} // 创建数据库

	var db *gorm.DB
	if db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       c.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: true,    // 根据版本自动配置
	}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return ctx, err
	}
	next = context.WithValue(next, "db", db)
	return next, err
}

func (h MySQLInitHandler) InitTables(ctx context.Context, inits initSlice) error {
	return createTables(ctx, inits)
}

func (h MySQLInitHandler) InitData(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for _, init := range inits {
		if init.DataInserted(next) {
			color.Info.Printf(InitDataExist, MySQL, init.InitializerName())
			continue
		}
		if n, err := init.InitializeData(next); err != nil {
			color.Info.Printf(InitDataFailed, MySQL, init.InitializerName(), err)
			return err
		} else {
			next = n
			color.Info.Printf(InitDataSuccess, MySQL, init.InitializerName())
		}
	}
	color.Info.Printf(InitSuccess, MySQL)
	return nil
}
