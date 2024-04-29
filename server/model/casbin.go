package model

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zilanlann/acmer-manage-system/server/global"
)


func CasbinSetup() {
	a, _ := gormadapter.NewAdapterByDB(db)
	global.Casbin, _ = casbin.NewEnforcer("./conf/model.conf", a)

	global.Casbin.LoadPolicy()
}
