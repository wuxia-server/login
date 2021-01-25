package Code

import (
	"github.com/wuxia-server/login/Cmd"
)

/**
 * 账户注册
 * 规则: (CMD*100) + iota
 * 例如: CMD=1001; CODE=100101
 */
const (
	_ = Cmd.Account_Register*100 + iota

	Account_Register_AlreadyExists         // 账户已存在
	Account_Register_GenerateAccountIdFail // AccountId生成失败 (AccountId生成次数达上限)
	Account_Register_GenerateTokenFail     // Token生成失败 (Token生成次数达上限)
)

/**
 * 账户密码登录
 * 规则: (CMD*100) + iota
 * 例如: CMD=1001; CODE=100101
 */
const (
	_ = Cmd.Account_PasswordLogin*100 + iota

	Account_PasswordLogin_NotExists         // 账户不存在
	Account_PasswordLogin_PasswordIncorrect // 密码错误
	Account_PasswordLogin_RefreshTokenFail  // 刷新Token失败
)

/**
 * 账户Token登录
 * 规则: (CMD*100) + iota
 * 例如: CMD=1001; CODE=100101
 */
const (
	_ = Cmd.Account_TokenLogin*100 + iota

	Account_TokenLogin_TokenIncorrect // Token错误
)

/**
 * 服务器列表
 * 规则: (CMD*100) + iota
 * 例如: CMD=1001; CODE=100101
 */
const (
	_ = Cmd.Server_List*100 + iota

	Server_List_TokenIncorrect // Token错误
)

/**
 * 选服
 * 规则: (CMD*100) + iota
 * 例如: CMD=1001; CODE=100101
 */
const (
	_ = Cmd.Server_Select*100 + iota

	Server_Select_TokenIncorrect // Token错误
	Server_Select_NotExists      // 该服不存在
	Server_Select_AccountDisable // 该账户已被禁用 (也许是被封号了)
)
