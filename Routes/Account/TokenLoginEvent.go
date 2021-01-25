package Account

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/logger"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/login/Code"
	"github.com/wuxia-server/login/Data"
	"net/http"
	"strings"
)

type TokenLoginEvent struct {
	Network.HttpRoute

	Token string // Token值
}

func (e *TokenLoginEvent) Parse() {
	e.Token = utils.NewStringAny(e.Params["token"]).ToString()
}

func (e *TokenLoginEvent) Handle(req *http.Request) uint32 {
	account := Data.GetAccountByToken(e.Token)

	// Token错误
	if account == nil {
		logger.Debug("Token错误")
		return Code.Account_TokenLogin_TokenIncorrect
	}

	// 账户信息
	e.Data("account", account.ToJsonMap())

	// 默认选中的服务器
	serverList := Data.GetServerList()
	if account.LatelyServer == "" {
		e.Data("server", serverList[0].ToJsonMap())
	} else {
		serverId := strings.Split(account.LatelyServer, ",")[0]
		for _, server := range serverList {
			if string(server.Id) == serverId {
				e.Data("server", serverList[0].ToJsonMap())
				break
			}
		}
	}
	return messages.RC_Success
}
