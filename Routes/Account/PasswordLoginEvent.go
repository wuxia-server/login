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

type PasswordLoginEvent struct {
	Network.HttpRoute

	UserName string // 账号
	PassWord string // 密码
}

func (e *PasswordLoginEvent) Parse() {
	e.UserName = utils.NewStringAny(e.Params["username"]).ToString()
	e.PassWord = utils.NewStringAny(e.Params["password"]).ToString()
}

func (e *PasswordLoginEvent) Handle(req *http.Request) uint32 {
	account := Data.GetAccountByUserName(e.UserName)

	// 账户不存在
	if account == nil {
		logger.Debug("账户不存在")
		return Code.Account_PasswordLogin_NotExists
	}

	// 密码错误
	if account.PassWord != e.PassWord {
		logger.Debug("密码错误")
		return Code.Account_PasswordLogin_PasswordIncorrect
	}

	// 刷新Token失败
	if !Data.RefreshAccountToken(account) {
		logger.Debug("刷新Token失败")
		return Code.Account_PasswordLogin_RefreshTokenFail
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
