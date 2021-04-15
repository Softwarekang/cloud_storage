package SQL

import "github.com/xormplus/xorm"

// 封装SQL命令
type SQLClient interface {
	SQL(db interface{}, query interface{}, args ...interface{}) *xorm.Session
	Insert(db interface{}, beans ...interface{}) (int64, error)
	Get(db interface{}, bean interface{}) (bool, error)
	Where(db interface{}, query interface{}, args ...interface{}) *xorm.Session
	ID(db interface{}, id interface{}) *xorm.Session
}
