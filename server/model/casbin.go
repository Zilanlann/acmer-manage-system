package model

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

var Casbin *casbin.Enforcer

func CasbinSetup() {
	a, _ := gormadapter.NewAdapterByDB(db)
	Casbin, _ = casbin.NewEnforcer("./conf/model.conf", a)

	Casbin.LoadPolicy()

	// casbinEnforcer.AddPolicy("Super Admin", "data1", "read")
	// // Check the permission.
	// if ok, _ := casbinEnforcer.Enforce("Super Admin", "data1", "read"); ok {
	// 	fmt.Println("passed")
	// } else {
	// 	fmt.Println("not passed")
	// }
	// casbinEnforcer.RemovePolicy("Super Admin", "data1", "read")
}
