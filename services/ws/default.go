package ws

import (
	"context"
	"github.com/gorilla/websocket"
)

var userMap = make(map[int]*UserWs)
var connMap = make(map[*websocket.Conn]*ConnInfo)

type ws struct {
}

// 登录用户的websocket链接
type UserWs struct {
	Uid       int32
	ClientNum uint8
	Conns     []*websocket.Conn
}

// 连接信息
type ConnInfo struct {
	Uid           int32
	LatestMsgTime int64 //  最后一条消息时间
	CreateTime    int64 //  创建时间
	Ip            string
}

func NewWs() *ws {
	return &ws{}
}

func (w *ws) Add(conn *websocket.Conn, ip string) {
	ctx, cancelFunc := context.WithCancel(context.Background())

}
