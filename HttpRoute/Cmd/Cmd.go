package Cmd

/**
 * 账户模块
 */
const (
	Account_Register      uint32 = 1001 // 注册
	Account_PasswordLogin uint32 = 1002 // 密码登录
	Account_TokenLogin    uint32 = 1003 // Token登录
)

/**
 * 服务器模块
 */
const (
	Server_List   uint32 = 2001 // 列表
	Server_Select uint32 = 2002 // 选服
)
