package ws

import (
	"github.com/gorilla/websocket"
	"sync"
)

// 普通map
type wsMap struct {
	m *sync.Map
}

func (wMap *wsMap) Delete(key *websocket.Conn) {
	wMap.m.Delete(key)
}

func (wMap *wsMap) Load(key *websocket.Conn) (value *ConnInfo, ok bool) {

	v, ok := wMap.m.Load(key)
	if v != nil {
		value = v.(*ConnInfo)
	}
	return
}

func (wMap *wsMap) LoadOrStore(key *websocket.Conn, value *ConnInfo) (actual *ConnInfo, loaded bool) {
	a, loaded := wMap.m.LoadOrStore(key, value)
	actual = a.(*ConnInfo)
	return
}

func (wMap *wsMap) Range(f func(key *websocket.Conn, value *ConnInfo) bool) {
	f1 := func(key, value interface{}) bool {
		return f(key.(*websocket.Conn), value.(*ConnInfo))
	}
	wMap.m.Range(f1)
}

func (wMap *wsMap) Store(key *websocket.Conn, value *ConnInfo) {
	wMap.m.Store(key, value)
}

// 用户map
type wsUserMap struct {
	m *sync.Map
}

func (wMap *wsUserMap) Delete(key int32) {
	wMap.m.Delete(key)
}

func (wMap *wsUserMap) Load(key int32) (value *UserWs, ok bool) {
	v, ok := wMap.m.Load(key)
	if v != nil {
		value = v.(*UserWs)
	}
	return
}

func (wMap *wsUserMap) LoadOrStore(key int32, value *UserWs) (actual *UserWs, loaded bool) {
	a, loaded := wMap.m.LoadOrStore(key, value)
	actual = a.(*UserWs)
	return
}

func (wMap *wsUserMap) Range(f func(key int32, value *UserWs) bool) {
	f1 := func(key, value interface{}) bool {
		return f(key.(int32), value.(*UserWs))
	}
	wMap.m.Range(f1)
}

func (wMap *wsUserMap) Store(key int32, value *UserWs) {
	wMap.m.Store(key, value)
}

type wsLockMap struct {
	m *sync.Map
}

func (ck *wsLockMap) Delete(key *websocket.Conn) {
	ck.m.Delete(key)
}

func (ck *wsLockMap) Load(key *websocket.Conn) (value *sync.Mutex, ok bool) {
	v, ok := ck.m.Load(key)
	if v != nil {
		value = v.(*sync.Mutex)
	}
	return
}

func (ck *wsLockMap) LoadOrStore(key *websocket.Conn, value *sync.Mutex) (actual *sync.Mutex, loaded bool) {
	a, loaded := ck.m.LoadOrStore(key, value)
	actual = a.(*sync.Mutex)
	return
}

func (ck *wsLockMap) Range(f func(key *websocket.Conn, value *sync.Mutex) bool) {
	f1 := func(key, value interface{}) bool {
		return f(key.(*websocket.Conn), value.(*sync.Mutex))
	}
	ck.m.Range(f1)
}

func (ck *wsLockMap) Store(key *websocket.Conn, value *sync.Mutex) {
	ck.m.Store(key, value)
}
