package HttpRoute

import (
	"github.com/team-zf/framework/Network"
	"github.com/wuxia-server/login/Cmd"
	"github.com/wuxia-server/login/HttpRoute/Account"
	"github.com/wuxia-server/login/HttpRoute/Server"
)

var (
	Route = Network.NewHttpRouteHandle()
)

func init() {
	Route.SetRoute(Cmd.Account_Register, &Account.RegisterEvent{})
	Route.SetRoute(Cmd.Account_PasswordLogin, &Account.PasswordLoginEvent{})
	Route.SetRoute(Cmd.Account_TokenLogin, &Account.TokenLoginEvent{})
	Route.SetRoute(Cmd.Server_List, &Server.ListEvent{})
	Route.SetRoute(Cmd.Server_Select, &Server.SelectEvent{})
}
