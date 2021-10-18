package code

// Ok 成功
const Ok int = 0

//==================客户端 400xxx==================
// ParamsIsInvalid 参数无效
const ParamsIsInvalid int = 400001

// LoginFailed 登录失败
const LoginFailed int = 400002

// UserStatusDisable 用户状态不可用
const UserStatusDisable int = 400003

//==================服务端 500xxx==================
// ServerError 服务端错误
const ServerError int = 500001

// ServerError 数据库错误
const DatabaseError int = 500002
const DatabaseRowsAffectedError int = 500003

// UserInfoFailed 用户信息错误
const UserInfoFailed int = 500004

// UserRegisterFailed 用户注册失败
const UserRegisterFailed int = 500005
