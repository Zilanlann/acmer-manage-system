package model

import (
	"fmt"

	"github.com/zilanlann/acmer-manage-system/server/global"
	"github.com/zilanlann/acmer-manage-system/server/setting"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Setup() {
	var err error

	switch setting.DatabaseSetting.Type {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", setting.DatabaseSetting.User, setting.DatabaseSetting.Password, setting.DatabaseSetting.Host, setting.DatabaseSetting.Port, setting.DatabaseSetting.Name)
		global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   setting.DatabaseSetting.TablePrefix, // table name prefix, table for `User` would be `t_users`
				SingularTable: true,                                // use singular table name, table for `User` would be `user` with this option enabled
			},
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			global.LOG.Error(err.Error())
		}
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Shanghai", setting.DatabaseSetting.Host, setting.DatabaseSetting.User, setting.DatabaseSetting.Password, setting.DatabaseSetting.Name, setting.DatabaseSetting.Port)
		global.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   setting.DatabaseSetting.TablePrefix, // table name prefix, table for `User` would be `t_users`
				SingularTable: true,                                // use singular table name, table for `User` would be `user` with this option enabled
			},
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			global.LOG.Error(err.Error())
		}
	}

	sqlDB, err := global.DB.DB()
	if err != nil {
		global.LOG.Error(err.Error())
	}
	sqlDB.SetMaxIdleConns(setting.DatabaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(setting.DatabaseSetting.MaxOpenConns)

	if err = global.DB.AutoMigrate(&User{}, &Contest{}, &Team{}, &Contestant{}, &OJContest{}, &Site{}, &SiteType{}); err != nil {
		global.LOG.Error(err.Error())
	}
	if err := CreateAdmin(); err != nil {
		global.LOG.Error(err.Error())
	}

	CasbinSetup()
}
