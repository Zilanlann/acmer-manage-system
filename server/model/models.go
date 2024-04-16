package model

import (
	"fmt"
	"log"

	"github.com/zilanlann/acmer-manage-system/server/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func Setup() {
	var err error

	switch setting.DatabaseSetting.Type {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", setting.DatabaseSetting.User, setting.DatabaseSetting.Password, setting.DatabaseSetting.Host, setting.DatabaseSetting.Port, setting.DatabaseSetting.Name)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   setting.DatabaseSetting.TablePrefix, // table name prefix, table for `User` would be `t_users`
				SingularTable: true,                                // use singular table name, table for `User` would be `user` with this option enabled
			},
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			log.Println(err)
		}
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Shanghai", setting.DatabaseSetting.Host, setting.DatabaseSetting.User, setting.DatabaseSetting.Password, setting.DatabaseSetting.Name, setting.DatabaseSetting.Port)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   setting.DatabaseSetting.TablePrefix, // table name prefix, table for `User` would be `t_users`
				SingularTable: true,                                // use singular table name, table for `User` would be `user` with this option enabled
			},
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			log.Println(err)
		}
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Println(err)
	}
	sqlDB.SetMaxIdleConns(setting.DatabaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(setting.DatabaseSetting.MaxOpenConns)

	db.AutoMigrate(&User{})
	AddAdmin()

	CasbinSetup()
}
