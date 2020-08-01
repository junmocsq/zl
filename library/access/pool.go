package access

import (
	"github.com/astaxie/beego"
	"github.com/silenceper/pool"
	"time"
)

type accessor struct {
	pool pool.Pool
}

func registerAccessor() {
	//factory 创建连接的方法
	factory := func() (interface{}, error) { return &dao{}, nil }

	//close 关闭连接的方法
	close := func(v interface{}) error {
		v = nil
		return nil
	}

	//ping 检测连接的方法
	//ping := func(v interface{}) error { return nil }

	initCap := beego.AppConfig.DefaultInt("access.pool.init", 5)
	maxIdle := beego.AppConfig.DefaultInt("mysql.pool.maxidle", 5)
	maxCap := beego.AppConfig.DefaultInt("mysql.pool.maxcap", 5)
	//创建一个连接池： 初始化5，最大空闲连接是20，最大并发连接30
	poolConfig := &pool.Config{
		InitialCap: initCap, //资源池初始连接数
		MaxIdle:    maxIdle, //最大空闲连接数
		MaxCap:     maxCap,  //最大并发连接数
		Factory:    factory,
		Close:      close,
		//Ping:       ping,
		//连接最大空闲时间，超过该时间的连接 将会关闭，可避免空闲时连接EOF，自动失效的问题
		IdleTimeout: 5 * time.Second,
	}
	var err error
	accessPool.pool, err = pool.NewChannelPool(poolConfig)
	if err != nil {
		panic("accessor pool init failed！")
	}
}

func NewAccessor() accessor {
	return accessPool
}

func (p *accessor) Get() (Accessor, error) {
	v, err := p.pool.Get()
	if err != nil {
		return nil, err
	}
	return v.(Accessor), nil
}

func (p *accessor) Close(accessor Accessor) {
	p.pool.Put(accessor)
}

func (p *accessor) Release() {
	p.pool.Release()
}

func (p *accessor) Length() int {
	current := p.pool.Len()
	return current
}
