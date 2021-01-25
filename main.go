package main

import (
	"github.com/team-zf/framework"
	"github.com/team-zf/framework/DB"
	"github.com/team-zf/framework/Network"
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
		modules.AppSetPStatusTime(10*time.Second),
	)
	Control.App.OnConfigurationLoaded(func(app modules.IApp, conf *config.AppConfig) {
		// 载入数据库模块(账户服)
		if item := conf.Settings["gate_db"]; item != nil {
			settings := item.(map[string]interface{})
			Control.GateDB = DB.NewDataBaseModule(
				DB.DataBaseSetName("GateDB"),
				DB.DataBaseSetDsn(settings["dsn"].(string)),
			)
			app.AddModule(Control.GateDB)
		}

		// 载入HTTP服务模块
		app.AddModule(Network.NewHttpModule(
			Network.HttpSetName("账户服"),
			Network.HttpSetIpPort(":20300"),
			Network.HttpSetRoute(HttpRoute.Route),
		))
	})
	Control.App.Init()
	Control.App.Run()
}
