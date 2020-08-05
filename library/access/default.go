package access

// 数据访问入口
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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
