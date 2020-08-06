package ws

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type wsConnPool struct {
	clients     *wsMap
	userClients *wsUserMap
	connLocks   *wsLockMap
}

var connPool = &wsConnPool{
	clients: &wsMap{
		m: new(sync.Map),
	},
	userClients: &wsUserMap{
		m: new(sync.Map),
	},
	connLocks: &wsLockMap{
		m: new(sync.Map),
	},
}

type wsConn struct {
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

func NewWs() *wsConn {
	return &wsConn{}
}

func (w *wsConn) Conn(conn *websocket.Conn, ip string) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	connInfo := &ConnInfo{
		Uid:           0,
		LatestMsgTime: 0,
		CreateTime:    time.Now().Unix(),
		Ip:            ip,
	}
	connPool.clients.Store(conn, connInfo)
	go w.checkConnValid(conn, ctx)

	newHandle().MsgHandle(conn, cancelFunc)
}

func (w *wsConn) Close(conn *websocket.Conn) *ConnInfo {
	connInfo, ok := connPool.clients.Load(conn)
	conn.Close()
	// TODO 清除clients里的ConnInfo
	// TODO 清除userClients里的conn
	// TODO 清除connLocks里的conn
	if !ok {
		return nil
	}
	return connInfo
}

// 连接有效性检查
// 定时心跳检测
// 新连接60s内无交互消息关闭 之后360s无交互关闭
func (w *wsConn) checkConnValid(conn *websocket.Conn, ctx context.Context) {
	writeWait := 10 * time.Second
	pongWait := 72 * time.Second
	pingPeriod := (pongWait * 9) / 10
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			// 心跳检测
			if err := conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(writeWait)); err != nil {
				connInfo := w.Close(conn)
				if connInfo != nil {
					logrus.WithField("websocket", "ping").
						Infof("uid:%d addr:%s",
							connInfo.Uid, connInfo.Ip)
				}
				return
			}
			currTime := time.Now().Unix()
			v, ok := connPool.clients.Load(conn)
			if ok {
				if v.LatestMsgTime == 0 {
					if currTime-v.CreateTime > 60 {
						// 新连接由于未在60s内发送任何消息将被动关闭
						connInfo := w.Close(conn)
						if connInfo != nil {
							logrus.WithField("websocket", "autoClose").
								Infof("uid:%d  addr:%s createTime:%s 新连接由于未在60s内发送任何消息而被动关闭",
									connInfo.Uid, connInfo.Ip, time2Read(connInfo.CreateTime))
							return
						}
					}
				} else {
					if currTime-v.LatestMsgTime > 360 {
						connInfo := w.Close(conn)
						if connInfo != nil {
							// 在360s内无消息关闭
							logrus.WithField("websocket", "autoClose").
								Infof("uid:%d addr:%s createTime:%s latestTime:%s 在360s内无消息关闭",
									connInfo.Uid, connInfo.Ip, time2Read(connInfo.CreateTime), time2Read(connInfo.LatestMsgTime))
							return
						}
					}
				}
			}
		case <-ctx.Done():
			connInfo := w.Close(conn)
			if connInfo == nil {
				logrus.WithField("websocket", "ping").
					Infof("由context关闭 addr:%s", conn.RemoteAddr().String())
			} else {
				logrus.WithField("websocket", "ping").
					Infof("由context关闭 ip:%s uid:%d", connInfo.Ip, connInfo.Uid)
			}
			return
		}
	}
}

func time2Read(t int64) string {
	if t%86400 == 57600 {
		return time.Unix(t, 0).Format("2006-01-02")
	}
	return time.Unix(t, 0).Format("2006-01-02 15:04:05")
}
