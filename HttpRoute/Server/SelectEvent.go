package Server

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/logger"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/login/Code"
	"github.com/wuxia-server/login/Data"
	"net/http"
)

type SelectEvent struct {
	Network.HttpRoute

	Token    string // Token值
	ServerId int64  // 服务器ID
}

func (e *SelectEvent) Parse() {
	e.Token = utils.NewStringAny(e.Params["token"]).ToString()
	e.ServerId = utils.NewStringAny(e.Params["server_id"]).ToInt64V()
}

func (e *SelectEvent) Handle(req *http.Request) uint32 {
	account := Data.GetAccountByToken(e.Token)

	// Token错误
	if account == nil {
		logger.Debug("Token错误")
		return Code.Server_Select_TokenIncorrect
	}

	// 该账户已被禁用, 也许是被封了
	if account.Status == 1 {
		logger.Debug("该账户已被禁用, 也许是被封了")
		return Code.Server_Select_AccountDisable
	}

	server := Data.GetServerById(e.ServerId)

	// 该服不存在
	if server == nil {
		logger.Debug("该服不存在")
		return Code.Server_Select_NotExists
	}

	logger.Debug("请及时连接逻辑服.")

	e.Data("host", server.Host)
	e.Data("port", server.Port)
	return messages.RC_Success
}
