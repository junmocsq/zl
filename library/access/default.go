package access

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "wangqingshui/library/cache/redis" // 注册access使用的cache
)

const (
	MYSQL_WQS = "default"
)

var masterDbMap map[string]*sql.DB = make(map[string]*sql.DB)
var slaveDbMap map[string][]*sql.DB = make(map[string][]*sql.DB)
var accessPool accessor

func init() {
	// set default database
	registerAccessor()
}
