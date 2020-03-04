package models

import (
	"fmt"
	"ginAdmin/pkg/setting"
	"github.com/casbin"
	"github.com/gorm-adapter"
)

var CasbinEnforcer *casbin.Enforcer

func initCasbinEnforcer()  {

	var err error
	// 将数据库连接同步给插件， 插件用来操作数据库
	po,_ := gormadapter.NewAdapterByDB(db)

	CasbinEnforcer,err = casbin.NewEnforcer(setting.AppConfigPath +"rbac_model.conf", po)


	if err!=nil {
		fmt.Println(err.Error())
	}

	// 开启权限认证日志
	CasbinEnforcer.EnableLog(true)
	// 加载数据库中的策略
	err = CasbinEnforcer.LoadPolicy()

	if err!=nil {
		fmt.Println(err.Error())
	}
}



//	// admin 这个角色可以访问GET 方式访问 /api/v2/ping
//	res,_ := e.AddPolicy("admin","/api/v2/ping","GET")
// 将 test 用户加入一个角色中
//e.AddRoleForUser("test","root")
//e.AddRoleForUser("tom","admin")
// 请看规则中如果用户名为 root 则不受限制