package code

// meaning code Corresponding message
var meaning = map[int]string{
	Ok: "SUCCESS",
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

// Value get msg by code
func Value(code int) string {
	if msg, ok := meaning[code]; ok {
		return msg
	} else {
		return ""
	}
}
