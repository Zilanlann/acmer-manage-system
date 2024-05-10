package model

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zilanlann/acmer-manage-system/server/global"
)

func CasbinSetup() {
	var err error
	a, err := gormadapter.NewAdapterByDB(global.DB)
	if err != nil {
		global.LOG.Panic(err.Error())
	}
	global.Casbin, err = casbin.NewEnforcer("./conf/model.conf", a)
	if err != nil {
		global.LOG.Panic(err.Error())
	}

	global.Casbin.LoadPolicy()
}
