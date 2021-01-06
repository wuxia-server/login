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

type PasswordLoginEvent struct {
	messages.HttpMessage

	UserName string // 账号
	PassWord string // 密码
}

func (e *PasswordLoginEvent) Parse() {
	e.UserName = utils.NewStringAny(e.Params["username"]).ToString()
	e.PassWord = utils.NewStringAny(e.Params["password"]).ToString()
}

func (e *PasswordLoginEvent) HttpDirectCall(req *http.Request, resp *messages.HttpResponse) {
	account := Data.GetAccountByUserName(e.UserName)

	// 账户不存在
	if account == nil {
		logger.Debug("账户不存在")
		resp.Code = Code.Account_PasswordLogin_NotExists
		return
	}

	// 密码错误
	if account.PassWord != e.PassWord {
		logger.Debug("密码错误")
		resp.Code = Code.Account_PasswordLogin_PasswordIncorrect
		return
	}

	// 刷新Token失败
	if !Data.RefreshAccountToken(account) {
		logger.Debug("刷新Token失败")
		resp.Code = Code.Account_PasswordLogin_RefreshTokenFail
		return
	}

	logger.Debug("登录成功, Token更新为: %s", account.Token)

	// 账户信息
	resp.Data["account"] = account.ToJsonMap()

	// 默认选中的服务器
	serverList := Data.GetServerList()
	if account.LatelyServer == "" {
		resp.Data["server"] = serverList[0].ToJsonMap()
	} else {
		serverId := strings.Split(account.LatelyServer, ",")[0]
		for _, server := range serverList {
			if string(server.Id) == serverId {
				resp.Data["server"] = serverList[0].ToJsonMap()
				break
			}
		}
	}
	resp.Code = messages.RC_Success
}

func M_PasswordLogin() *PasswordLoginEvent {
	return &PasswordLoginEvent{}
}
