package Account

import (
	"github.com/team-zf/framework/logger"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/login/Data"
	"github.com/wuxia-server/login/HttpRoute/Code"
	"github.com/wuxia-server/login/Table"
	"net/http"
)

type RegisterEvent struct {
	messages.HttpMessage

	UserName string // 账号
	PassWord string // 密码
}

func (e *RegisterEvent) Parse() {
	e.UserName = utils.NewStringAny(e.Params["username"]).ToString()
	e.PassWord = utils.NewStringAny(e.Params["password"]).ToString()
}

func (e *RegisterEvent) HttpDirectCall(req *http.Request, resp *messages.HttpResponse) {
	// 账户已存在
	if Data.GetAccountByUserName(e.UserName) != nil {
		logger.Debug("账户已存在")
		resp.Code = Code.Account_Register_AlreadyExists
		return
	}

	var id int64
	var token string
	var ok bool

	// 获取并判定AccountId生成是否成功
	id, ok = Data.GenerateAccountId()
	if !ok {
		logger.Debug("AccountId生成失败")
		resp.Code = Code.Account_Register_GenerateAccountIdFail
		return
	}

	// 获取并判定Token生成是否成功
	token, ok = Data.GenerateToken()
	if !ok {
		logger.Debug("Token生成失败")
		resp.Code = Code.Account_Register_GenerateTokenFail
		return
	}

	account := Table.NewAccount()
	account.Id = id
	account.UserName = e.UserName
	account.PassWord = e.PassWord
	account.Token = token

	if err := Data.RegisterAccount(account); err != nil {
		logger.Error("注册失败, 原因: %+v", err)
		panic(err)
	}

	logger.Debug("注册成功")
	resp.Code = messages.RC_Success
	return
}

func M_Register() *RegisterEvent {
	return &RegisterEvent{}
}
