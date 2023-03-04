package code

const (
	OK = 0 //@msg 成功
	// ==================客户端 400xxx==================
	ParamsIsInvalid   = 400001 //@msg 参数无效
	LoginFailed       = 400002 //@msg 登录失败
	UserStatusDisable = 400003 //@msg 用户状态不可用

	// ==================服务端 500xxx==================
	ServerError               = 500001 //@msg 服务端错误
	DatabaseError             = 500002 //@msg 数据库错误
	DatabaseRowsAffectedError = 500003 //@msg 影响行数为零

	UserInfoFailed     = 500004 //@msg 用户信息错误
	UserRegisterFailed = 500005 //@msg 用户注册失败
)

// Message returns the correct message
func Message(code int) string {
	if msg, ok := statusMessage[code]; ok {
		return msg
	} else {
		return ""
	}
}
