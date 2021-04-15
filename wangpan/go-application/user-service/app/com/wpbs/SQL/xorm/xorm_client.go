package xorm

import (
	"github.com/xormplus/xorm"
	"user-service/app/com/wpbs/SQL"
	"user-service/common/extension"
)

func init() {
	extension.SetSQlClient("xorm", NewXormClient)
}

func NewXormClient() SQL.SQLClient {
	return &XormClient{}
}

type XormClient struct {
}

func (x *XormClient) SQL(db interface{}, query interface{}, args ...interface{}) *xorm.Session {
	engine, ok := db.(*xorm.Engine)
	if ok {
		return engine.SQL(query, args...)
	}
	session := db.(*xorm.Session)
	return session.SQL(query, args...)
}

func (x *XormClient) Insert(db interface{}, beans ...interface{}) (int64, error) {
	engine, ok := db.(*xorm.Engine)
	if ok {
		return engine.Insert(beans...)
	}
	session := db.(*xorm.Session)
	return session.Insert(beans...)
}

func (x *XormClient) Get(db interface{}, bean interface{}) (bool, error) {
	engine, ok := db.(*xorm.Engine)
	if ok {
		return engine.Get(bean)
	}
	session := db.(*xorm.Session)
	return session.Get(bean)
}
func (x *XormClient) Where(db interface{}, query interface{}, args ...interface{}) *xorm.Session {
	engine, ok := db.(*xorm.Engine)
	if ok {
		return engine.Where(query, args...)
	}
	session := db.(*xorm.Session)
	return session.Where(query, args...)
}

func (x *XormClient) ID(db interface{}, id interface{}) *xorm.Session {
	engine, ok := db.(*xorm.Engine)
	if ok {
		return engine.ID(id)
	}
	session := db.(*xorm.Session)
	return session.ID(id)
}
