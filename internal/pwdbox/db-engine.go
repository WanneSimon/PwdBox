package pwdbox

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/zhuxiujia/GoMybatis"
)

type Service interface {
	Init(sqlEngine GoMybatis.GoMybatisEngine) // 初始化
}

// var uri string = "file:D:\\eclipse\\2022-06-R\\workspace\\pwdbox.db3"
var SqlEngine GoMybatis.GoMybatisEngine
var SqliteInited bool = false // 数据库是否初始化成功

// 初始化数据库链接
func InitSqlite(uri string) {
	// uri = uri
	SqlEngine = GoMybatis.GoMybatisEngine{}.New()
	_, err := SqlEngine.Open("sqlite3", uri)
	if err != nil {
		panic(err)
	}
	SqliteInited = true

	// 初始化各个 service mapper
	GetPlatformService()
	GetAccountService()
}

// 接口操作
type DbOp struct {
}

// Platform 的service
func (do *DbOp) IsInited() bool {
	return SqliteInited
}

// Account的service
// func (do *DbOp) AccountOp() AccountService {
// 	return GetAccountService()
// }

// Platform 的service
// func (do *DbOp) PlatformOp() PlatformService {
// 	return GetPlatformService()
// }
