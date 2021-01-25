package Routes

import (
	"github.com/team-zf/framework/Network"
	"github.com/wuxia-server/login/Cmd"
	"github.com/wuxia-server/login/Routes/Account"
	"github.com/wuxia-server/login/Routes/Server"
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
