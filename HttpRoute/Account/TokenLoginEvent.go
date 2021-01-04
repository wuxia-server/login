package Account

import (
	"github.com/team-zf/framework/logger"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/login/Data"
	"github.com/wuxia-server/login/HttpRoute/Code"
	"net/http"
	"strings"
)

type TokenLoginEvent struct {
	messages.HttpMessage

	Token string // Token值
}

func (e *TokenLoginEvent) Parse() {
	e.Token = utils.NewStringAny(e.Params["token"]).ToString()
}

func (e *TokenLoginEvent) HttpDirectCall(req *http.Request, resp *messages.HttpResponse) {
	account := Data.GetAccountByToken(e.Token)

	// Token错误
	if account == nil {
		logger.Debug("Token错误")
		resp.Code = Code.Account_TokenLogin_TokenIncorrect
		return
	}

	logger.Debug("Token登录成功")

	// 账户信息
	resp.Data["Account"] = account.ToJsonMap()

	// 默认选中的服务器
	serverList := Data.GetServerList()
	if account.LatelyServer == "" {
		resp.Data["Server"] = serverList[0].ToJsonMap()
	} else {
		serverId := strings.Split(account.LatelyServer, ",")[0]
		for _, server := range serverList {
			if string(server.Id) == serverId {
				resp.Data["Server"] = serverList[0].ToJsonMap()
				break
			}
		}
	}
	resp.Code = messages.RC_Success
}

func M_TokenLogin() *TokenLoginEvent {
	return &TokenLoginEvent{}
}
