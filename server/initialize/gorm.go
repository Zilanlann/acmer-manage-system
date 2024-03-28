package initialize

import (
	"os"

	"github.com/zilanlann/acmer-manage-system/server/global"
	// "github.com/flipped-aurora/gin-vue-admin/server/model/example"
	// "github.com/flipped-aurora/gin-vue-admin/server/model/system"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormMySQL 初始化Mysql数据库
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func GormMySQL() *gorm.DB {
	m := global.AMS_CONFIG.MySQL
	if m.DBName == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// GormMysqlByConfig 初始化Mysql数据库用过传入配置
func GormMysqlByConfig(m config.Mysql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

func RegisterTables() {
	db := global.AMS_DB
	err := db.AutoMigrate(

	// system.SysApi{},
	// system.SysUser{},
	// system.SysBaseMenu{},
	// system.JwtBlacklist{},
	// system.SysAuthority{},
	// system.SysDictionary{},
	// system.SysOperationRecord{},
	// system.SysAutoCodeHistory{},
	// system.SysDictionaryDetail{},
	// system.SysBaseMenuParameter{},
	// system.SysBaseMenuBtn{},
	// system.SysAuthorityBtn{},
	// system.SysAutoCode{},
	// system.SysExportTemplate{},
	// system.Condition{},

	// example.ExaFile{},
	// example.ExaCustomer{},
	// example.ExaFileChunk{},
	// example.ExaFileUploadAndDownload{},
	)
	if err != nil {
		global.AMS_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.AMS_LOG.Info("register table success")
}
