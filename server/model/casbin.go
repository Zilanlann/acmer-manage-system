package model

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

var casbinEnforcer *casbin.Enforcer

func CasbinSetup() {
	a, _ := gormadapter.NewAdapterByDB(db)
	casbinEnforcer, _ = casbin.NewEnforcer("./conf/model.conf", a)

	casbinEnforcer.LoadPolicy()

	casbinEnforcer.AddPolicy("Super Admin", "data1", "read")
	// Check the permission.
	if ok, _ := casbinEnforcer.Enforce("Super Admin", "data1", "read"); ok {
		fmt.Println("passed")
	} else {
		fmt.Println("not passed")
	}
	casbinEnforcer.RemovePolicy("Super Admin", "data1", "read")

	casbinEnforcer.SavePolicy()
}
