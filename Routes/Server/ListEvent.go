package Server

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/logger"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/login/Code"
	"github.com/wuxia-server/login/Data"
	"github.com/wuxia-server/login/DataTable"
	"net/http"
	"strings"
)

type ListEvent struct {
	Network.HttpRoute

	Token string // Token值
}

func (e *ListEvent) Parse() {
	e.Token = utils.NewStringAny(e.Params["token"]).ToString()
}

func (e *ListEvent) Handle(req *http.Request) uint32 {
	account := Data.GetAccountByToken(e.Token)

	// Token错误
	if account == nil {
		logger.Debug("Token错误")
		return Code.Server_List_TokenIncorrect
	}

	serverList := Data.GetServerList()
	serverIds := strings.Split(account.LatelyServer, ",")

	latelyList := make([]*DataTable.Server, 0)
	for _, sid := range serverIds {
		for _, server := range serverList {
			if string(server.Id) == sid {
				latelyList = append(latelyList, server)
				break
			}
		}
	}

	// 最近登录的服务器列表
	if len(latelyList) > 0 {
		mlist := make([]map[string]interface{}, 0)
		for _, item := range latelyList {
			mlist = append(mlist, item.ToJsonMap())
		}
		e.Data("lately", mlist)
	}
	// 所有服务器列表
	if len(serverList) > 0 {
		mlist := make([]map[string]interface{}, 0)
		for _, item := range serverList {
			mlist = append(mlist, item.ToJsonMap())
		}
		e.Data("list", mlist)
	}

	return messages.RC_Success
}
