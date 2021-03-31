package store

import (
	"github.com/xormplus/xorm"
	"user-service/app/com/wpbs/dao"
)

var DBClient Store

func init() {
	DBClient = InitStore()
}

type Store interface {
	// 开启事务
	BeginTx() (*xorm.Session, error)
	// 关闭事务
	EndTx(session *xorm.Session, err error)

	// 返回非事务
	Begin() (engine *xorm.Engine)
	// 获取userDao
	User(DB interface{}, arg ...string) *dao.UserDao
	// 获取fileDao
	File(DB interface{}, arg ...string) *dao.FileDao
	// 获取 memorydao
	Memory(Db interface{}, arg ...string) *dao.MemoryDao
}
