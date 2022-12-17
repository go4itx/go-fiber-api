# fiber-rest-server 单仓库多应用、自动代码生成、swagger接口文档
基于golang、fiber、gorm的rest脚手架，快速开发web api

## 一、demo架子
基础crud项目结构
### 1.登录接口
post: /v1/login
```
{
    "name":"test001",
    "password":"123456"
}
```

### 2.获取登录用户信息
get: /v1/user/
```
header参数： Bearer Token
```

## 二、IM即时通信
### 1.配置说明
#### 根据自己情况配置websocket最大连接数
```
[server.im] 
    maxConns = 100
```
#### 允许使用http接口发送信息的登录账号
```
[account.admin]
    name = "admin"
    password = "123456"
``` 

### 2.指令说明
```
// 错误code
USER_IS_ONLINE   uint8 = 1 // 用户已在线
EXCEED_MAX_CONNS uint8 = 2 // 超出最大连接限制
INCOMPLETE_INFO  uint8 = 3  //信息不完整
```
```
//99：注册，100：踢单人下线 120：心跳，121：一对一消息，200：踢一组下线 ，201：群组消息， 210：全网广播...
KICK       uint8 = 100
REGISTER   uint8 = 101
HEARTBEAT  uint8 = 120
P2P        uint8 = 121
KICK_GROUP uint8 = 200
GROUP      uint8 = 201
BROADCAST  uint8 = 210
```
```
type Message struct {
    CMD  uint8       `json:"cmd" validate:"required"`
    From string      `json:"from" validate:"required"` // 发送者即用户id，必须保证一个唯一
    To   string      `json:"to" validate:"required"`   // cmd==10x是表示用户id，cmd==20x是表示群gid
    Body interface{} `json:"body" validate:"required"` // 消息内容
}
```
### 3.websocket连接例子（JavaScript）
```
let ws = new WebSocket("ws://127.0.0.1:20105/v1/im");
ws.onopen = function (evt) {
    let login = {
        "cmd": 101,
        "from": id,
        "to": gid,
        "body": ""
    }

    conn.send(JSON.stringify(login))
}
```

### 4.通过http接口操作

#### 登录接口(其它接口需要 Bearer Token)
post: /v1/im/login
```
{
    "name":"admin",
    "password":"123456im"
}
```
#### 发送消息
post: /v1/im/sendMessage
```
{
    "from":"system",
    "to":"-",
    "cmd":210,
    "body":{"msg":"广播消息001"}
}
```

#### 踢下线
post: /v1/im/kick
```
{
    "from":"system",
    "to":"test001",
    "cmd":100,
    "body":"踢用户下线"
}
```

#### 在线用户
get: /v1/im/online
```
参数：id或gid
```