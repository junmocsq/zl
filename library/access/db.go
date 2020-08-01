package access

import (
	"database/sql"
	"github.com/astaxie/beego"
	"math/rand"
	"time"
)

func masterDb(aliasName string) *sql.DB {
	if db, ok := masterDbMap[aliasName]; ok {
		return db
	}
	panic(aliasName + " 主数据库不存在 请配置")
	return nil
}

func slaveDb(aliasName string) *sql.DB {
	if dbArr, ok := slaveDbMap[aliasName]; ok {
		length := len(dbArr)
		if length == 1 {
			return dbArr[0]
		} else {
			rand.Seed(int64(time.Now().Nanosecond()))
			return dbArr[rand.Intn(length)]
		}
	}
	panic(aliasName + " 从数据库不存在 请配置")
	return nil
}

func RegisterMasterDb(aliasName, dbConfig string) {
	config := beego.AppConfig.DefaultString(dbConfig, "root:123456@tcp(127.0.0.1:3306)/wqs?charset=utf8")
	db, err := sql.Open("mysql", config)
	if err != nil {
		panic("数据库配置 " + config + " 失败")
	}
	err = db.Ping()
	if err != nil {
		panic("数据库 " + config + " ping失败")
	}

	masterDbMap[aliasName] = db
}

func RegisterSlaveDb(aliasName, dbConfig string) {
	config := beego.AppConfig.DefaultString(dbConfig, "root:123456@tcp(127.0.0.1:3306)/wqs?charset=utf8")
	db, err := sql.Open("mysql", config)
	if err != nil {
		panic("数据库配置失败")
	}
	err = db.Ping()
	if err != nil {
		panic("数据库ping失败")
	}

	slaveDbMap[aliasName] = append(slaveDbMap[aliasName], db)

}

func RegisterSlaveDbSameMasterDb(aliasName string) {
	db := masterDb(aliasName)
	slaveDbMap[aliasName] = append(slaveDbMap[aliasName], db)
}
