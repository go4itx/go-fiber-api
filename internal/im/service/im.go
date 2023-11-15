package service

import (
	conf "home/config"
	"home/pkg/code"
	"home/pkg/e"
	"log"
	"sync"
	"time"

	"github.com/gofiber/websocket/v2"
)

const (
	FROM_SYSTEM string = "system"

	// 错误code
	USER_IS_ONLINE   uint8 = 1 // 用户已在线
	EXCEED_MAX_CONNS uint8 = 2 // 超出最大连接限制
	INCOMPLETE_INFO  uint8 = 3 //信息不完整

	//99：注册，100：踢单人下线 120：心跳，121：一对一消息，200：踢一组下线 ，201：群组消息， 210：全网广播...
	KICK       uint8 = 100
	REGISTER   uint8 = 101
	HEARTBEAT  uint8 = 120
	P2P        uint8 = 121
	KICK_GROUP uint8 = 200
	GROUP      uint8 = 201
	BROADCAST  uint8 = 210
)

type Message struct {
	CMD  uint8       `json:"cmd" validate:"required"`
	From string      `json:"from" validate:"required"` // 发送者即用户id，必须保证一个唯一
	To   string      `json:"to" validate:"required"`   // cmd==10x是表示用户id，cmd==20x是表示群gid
	Body interface{} `json:"body" validate:"required"` // 消息内容
}

type User struct {
	ID  string `json:"id"`
	GID string `json:"gid"`
}

type Client struct {
	user User
	conn *websocket.Conn
}

type Conn struct {
	sync.RWMutex
	clients map[string]Client
}

// set Conn
func (conn *Conn) set(id string, client Client) {
	conn.Lock()
	conn.clients[id] = client
	conn.Unlock()
}

// get Conn
func (conn *Conn) get(id string) (client Client, err error) {
	conn.RLock()
	client, ok := conn.clients[id]
	conn.RUnlock()
	if !ok {
		err = e.NewError(code.ParamsIsInvalid, "User not online")
		return
	}

	return
}

// list Conn
func (conn *Conn) list(gid string) (clients []Client) {
	conn.RLock()
	for _, client := range conn.clients {
		if gid != "" && gid != client.user.GID {
			continue
		}

		clients = append(clients, client)
	}

	conn.RUnlock()
	return
}

// delete Conn
func (conn *Conn) delete(id string) {
	if id != "" {
		conn.Lock()
		delete(conn.clients, id)
		conn.Unlock()
	}
}

type Send struct {
	client  Client
	message Message // 消息
}

type imService struct {
	conn    Conn
	count   chan struct{}
	send    chan Send
	message chan Message
}

// newIMService construct an instance
func newIMService() *imService {
	return &imService{
		send:    make(chan Send, 200),
		message: make(chan Message, 200),
		count:   make(chan struct{}, conf.Viper.GetInt("server.im.maxConns")),
		conn: Conn{
			RWMutex: sync.RWMutex{},
			clients: make(map[string]Client),
		},
	}
}

// Kick 踢用户下线
func (s *imService) Kick(message Message) (err error) {
	switch message.CMD {
	case KICK, KICK_GROUP: // 单人/组
		s.message <- message
	default:
		return e.NewError(code.ParamsIsInvalid)
	}

	return
}

// Online imService
func (s *imService) Online(user User) (res []User, count int, total map[string]interface{}, err error) {
	defer func() {
		count = len(res)
		total = map[string]interface{}{
			"total": len(s.count),
		}
	}()

	var client Client
	res = make([]User, 0)
	if user.ID != "" {
		client, err = s.conn.get(user.ID)
		res = append(res, client.user)
		return
	}

	for _, client := range s.conn.list(user.GID) {
		res = append(res, client.user)
	}

	return
}

// SendMessage imService
func (s *imService) SendMessage(message Message) (err error) {
	switch message.CMD {
	case P2P, GROUP, BROADCAST:
		s.message <- message
	default:
		err = e.NewError(code.ParamsIsInvalid)
	}

	return
}

// WS imService
func (s *imService) WS(conn *websocket.Conn) {
	var (
		message   Message
		client    = Client{conn: conn}
		connected = false
	)

	defer func() {
		errorHandler("Connect error")
		s.conn.delete(client.user.ID)
		if connected {
			<-s.count
		}

		if err := conn.Close(); err != nil {
			log.Println("Close error:", client.user.ID, err)
		}
	}()

	select {
	case s.count <- struct{}{}:
		connected = true
	case <-time.After(time.Second * 1):
		conn.WriteJSON(Message{
			CMD:  EXCEED_MAX_CONNS,
			From: FROM_SYSTEM,
			To:   client.user.ID,
			Body: "Connection failed: too many people, please wait",
		})

		return
	}

	for {
		if err := conn.ReadJSON(&message); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("ReadJSON error:", err)
			}

			return
		}

		switch message.CMD {
		case P2P, GROUP: // 一对一、群消息
			s.message <- message
		case HEARTBEAT: // 心跳
			conn.WriteJSON(Message{
				CMD:  HEARTBEAT,
				From: FROM_SYSTEM,
				To:   client.user.ID,
				Body: "Heartbeat",
			})
		case REGISTER: // 注册
			if message.From == "" || message.From == message.To {
				conn.WriteJSON(Message{
					CMD:  INCOMPLETE_INFO,
					From: FROM_SYSTEM,
					To:   client.user.ID,
					Body: "ID or GID can not be empty",
				})

				return
			}

			// 检查是否在线
			if _, err := s.conn.get(message.From); err == nil {
				conn.WriteJSON(Message{
					CMD:  USER_IS_ONLINE,
					From: FROM_SYSTEM,
					To:   client.user.ID,
					Body: "Connection failed: user is online",
				})

				return
			}

			client.user.ID = message.From
			client.user.GID = message.To
			s.conn.set(client.user.ID, client)
		}
	}
}

// receiver 接收器
func (s *imService) receiver() {
	defer errorHandler("receiver error")
	for {
		message := <-s.message
		if message.Body == nil || message.Body == "" {
			continue
		}

		switch message.CMD {
		case P2P, KICK:
			if client, err := s.conn.get(message.To); err == nil {
				s.send <- Send{client: client, message: message}
			}
		case GROUP, KICK_GROUP, BROADCAST:
			gid := message.To
			if message.CMD == BROADCAST {
				gid = ""
			}

			for _, client := range s.conn.list(gid) {
				s.send <- Send{client: client, message: message}
			}
		}
	}
}

// sender 发送器
func (s *imService) sender() {
	defer errorHandler("sender error")
	for {
		send := <-s.send
		if err := send.client.conn.WriteJSON(send.message); err != nil {
			send.client.conn.SetReadDeadline(time.Now())
		} else {
			switch send.message.CMD {
			case USER_IS_ONLINE, EXCEED_MAX_CONNS, KICK, KICK_GROUP:
				send.client.conn.SetReadDeadline(time.Now())
			}
		}
	}
}
