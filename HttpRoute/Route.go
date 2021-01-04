package HttpRoute

import (
	"github.com/team-zf/framework/messages"
	"github.com/wuxia-server/login/HttpRoute/Account"
	"github.com/wuxia-server/login/HttpRoute/Cmd"
	"github.com/wuxia-server/login/HttpRoute/Server"
)

var (
	Route = messages.NewHttpMessageHandle()
)

func init() {
	Route.SetRoute(Cmd.Account_Register, Account.M_Register())
	Route.SetRoute(Cmd.Account_PasswordLogin, Account.M_PasswordLogin())
	Route.SetRoute(Cmd.Account_TokenLogin, Account.M_TokenLogin())
	Route.SetRoute(Cmd.Server_List, Server.M_List())
	Route.SetRoute(Cmd.Server_Select, Server.M_Select())
}
