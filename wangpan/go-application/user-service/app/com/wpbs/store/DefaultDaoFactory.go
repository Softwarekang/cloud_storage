package store

import (
	"github.com/xormplus/xorm"
	"user-service/app/com/wpbs/config/database"
	"user-service/app/com/wpbs/dao"
	"user-service/common"
)

var (
	log = common.GetLogger()
)

type DefaultFactory struct {
	Store Store
	DB    *xorm.Engine
}

func InitStore() Store {
	defalutFactory := new(DefaultFactory)
	defalutFactory.DB = database.GetDB()
	return defalutFactory
}
func (d *DefaultFactory) BeginTx() (*xorm.Session, error) {
	session := d.DB.NewSession()
	err := session.Begin()
	if err != nil {
		log.Errorf("begin tx error:", err)
		return nil, err
	}

	return session, nil
}

func (d *DefaultFactory) EndTx(session *xorm.Session, err error) {
	if err != nil {
		session.Rollback()
		return
	}
	session.Commit()
}

func (d *DefaultFactory) Begin() (engine *xorm.Engine) {
	return d.DB
}

func (d *DefaultFactory) User(DB interface{}) *dao.UserDao {
	return dao.NewUserDao(DB)
}

func (d *DefaultFactory) File(DB interface{}) *dao.FileDao {
	return dao.NewFileDao(DB)
}
