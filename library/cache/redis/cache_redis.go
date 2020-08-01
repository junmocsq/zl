package redis

import (
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
	"time"
	"wangqingshui/library/cache"
)

var pool *redis.Pool

func init() {
	var c cache.Cache = &cacheRedis{}
	cache.RegisterCache(c)

	host := beego.AppConfig.DefaultString("cache.redis.host", "127.0.0.1")
	port := beego.AppConfig.DefaultString("cache.redis.port", "6379")
	auth := beego.AppConfig.DefaultString("cache.redis.auth", "")
	maxIdle := beego.AppConfig.DefaultInt("cache.redis.maxidle", 32)
	maxActive := beego.AppConfig.DefaultInt("cache.redis.maxactive", 1024)
	pool = initRedis(host, port, auth, maxIdle, maxActive)
}

func initRedis(host, port, auth string, maxIdle, maxActive int) *redis.Pool {
	pool := &redis.Pool{
		MaxIdle:     maxIdle,   // 最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
		MaxActive:   maxActive, // 最大的连接数，表示同时最多有N个连接。0表示不限制。
		IdleTimeout: time.Duration(120),
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				host+":"+port,
				redis.DialReadTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialDatabase(0),
				redis.DialPassword(auth),
			)
		},
	}
	return pool
}

func redisRegister() {

}

type cacheRedis struct {
}

func (r *cacheRedis) Get(key string) (string, error) {
	con := pool.Get()
	if err := con.Err(); err != nil {
		return "", err
	}
	defer con.Close()
	return redis.String(con.Do("GET", key))
}

func (r *cacheRedis) Set(key string, val string) (bool, error) {
	con := pool.Get()
	if err := con.Err(); err != nil {
		return false, err
	}
	defer con.Close()
	res, err := con.Do("SET", key, val)
	if err != nil {
		return false, err
	}
	if res == "OK" {
		return true, nil
	}
	return false, nil
}

func (r *cacheRedis) SetEx(key string, val string, expire int) (bool, error) {
	con := pool.Get()
	if err := con.Err(); err != nil {
		return false, err
	}
	defer con.Close()
	res, err := con.Do("SET", key, val, "EX", expire)
	if err != nil {
		return false, err
	}
	if res == "OK" {
		return true, nil
	}
	return false, nil
}

func (r *cacheRedis) Delete(key string) (bool, error) {
	con := pool.Get()
	if err := con.Err(); err != nil {
		return false, err
	}
	defer con.Close()
	return redis.Bool(con.Do("DEL", key))
}

func (r *cacheRedis) SetTimeout(key string, expire int) (bool, error) {
	con := pool.Get()
	if err := con.Err(); err != nil {
		return false, err
	}
	defer con.Close()
	return redis.Bool(con.Do("EXPIRE", key, expire))
}

func ttl(key string) (int, error) {
	con := pool.Get()
	if err := con.Err(); err != nil {
		return 0, err
	}
	defer con.Close()
	return redis.Int(con.Do("TTL", key))
}
