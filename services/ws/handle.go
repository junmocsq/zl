package ws

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"strings"
	"sync"
	"time"
)

type handler interface {
}

type handle struct {
}

func newHandle() *handle {
	return &handle{}
}

func (h *handle) MsgHandle(conn *websocket.Conn, cancelFunc context.CancelFunc) {
	defer func() {
		cancelFunc()
		err := recover()
		if err != nil {
			logrus.WithField("websocket", "MsgHandle").Errorf("panic 消息处理错误, err:%s", err)
		}
	}()
	var errNum int
	for {
		var msg Message
		err := conn.ReadJSON(&msg)

		if err != nil {
			logrus.WithField("MsgHandle", "MsgHandle").
				Errorf("读取信息失败，err:%s", err.Error())
			if _, ok := err.(*websocket.CloseError); ok {
				return
			}
			if strings.Contains(err.Error(), "use of closed network connection") {
				return
			}
			errNum++
			// 错误10次也关闭
			if errNum > 10 {
				return
			}
			continue
		} else {
			connInfo, ok := connPool.clients.Load(conn)
			ip := ""
			if ok {
				ip = connInfo.Ip
			}
			logrus.WithField("websocket", "MsgHandle").
				Debugf("读取信息内容 ipArr:%s，msg:%#v", ip, msg)
		}
		switch msg.Category {
		case CATE_HEART:
		case CATE_PRIVATE_MSG:
		case CATE_COMMENT:
		case CATE_TOPIC:
		case CATE_LOGIN:
		default:

		}
	}
}

func (h *handle) Heart(msg Message)      {}
func (h *handle) PrivateMsg(msg Message) {}
func (h *handle) Comment(msg Message)    {}
func (h *handle) Topic(msg Message)      {}
func (h *handle) Login(msg Message)      {}

// 数据推送方法
func (h *handle) writeMsg(conn *websocket.Conn, msg Message) bool {
	if msg.Ack == 1 {
		connInfo, ok := connPool.clients.Load(conn)
		if !ok {
			NewWs().Close(conn)
		}
		connInfo.LatestMsgTime = time.Now().Unix()
		connPool.clients.Store(conn, connInfo)
		return false
	}

	writePushMsgMutex, _ := connPool.connLocks.LoadOrStore(conn, new(sync.Mutex))
	writePushMsgMutex.Lock()
	defer writePushMsgMutex.Unlock()
	err := conn.WriteJSON(msg)
	if err != nil {
		logrus.WithField("websocket", "writeMsg").
			Errorf("MsgErr:%s", err.Error())
		NewWs().Close(conn)
		return false
	}
	return true
}
