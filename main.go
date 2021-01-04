package main

import (
	"github.com/team-zf/framework"
	"github.com/team-zf/framework/config"
	"github.com/team-zf/framework/modules"
	"github.com/wuxia-server/login/Control"
	"github.com/wuxia-server/login/HttpRoute"
	"time"
)

func main() {
	Control.App = framework.CreateApp(
		modules.AppSetDebug(true),
		modules.AppSetParse(true),
		modules.AppSetPStatusTime(3*time.Second),
	)
	Control.App.OnConfigurationLoaded(func(app modules.IApp, conf *config.AppConfig) {
		// 载入数据库模块
		Control.DbModule = modules.NewDataBaseModule(
			modules.DataBaseSetConf(conf.MySql),
		)
		app.AddModule(Control.DbModule)

		// 载入HTTP服务模块
		app.AddModule(modules.NewHttpModule(
			modules.HttpSetRoute(HttpRoute.Route),
		))
	})
	Control.App.Init()
	Control.App.Run()
}
