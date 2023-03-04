package code

// statusMessage NOTE: Keep this in sync with the status code list
var statusMessage = map[int]string{
	OK:                        "成功",
	ParamsIsInvalid:           "参数无效",
	LoginFailed:               "登录失败",
	UserStatusDisable:         "用户状态不可用",
	ServerError:               "服务端错误",
	DatabaseError:             "数据库错误",
	DatabaseRowsAffectedError: "影响行数为零",
	UserInfoFailed:            "用户信息错误",
	UserRegisterFailed:        "用户注册失败",
}
