package code

// codeMsg code Corresponding message
var codeMsg = map[int]string{
	OK: "SUCCESS",
	// client
	ParamsIsInvalid:   "参数无效",
	LoginFailed:       "账号或密码错误",
	UserStatusDisable: "账号已禁用",

	// server
	ServerError:               "服务端错误",
	DatabaseError:             "数据库错误",
	DatabaseRowsAffectedError: "影响行数为零",
	UserInfoFailed:            "获取用户失败",
	UserRegisterFailed:        "用户注册失败",
}

// Message get msg by code
func Message(code int) string {
	if msg, ok := codeMsg[code]; ok {
		return msg
	} else {
		return ""
	}
}
