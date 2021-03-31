package dao

import (
	"time"
	"user-service/app/com/wpbs/DTO"
	"user-service/app/com/wpbs/PO"
	"user-service/common/extension"
	log2 "user-service/common/log"
)

var (
	log = log2.GetLogger()
)

type UserDao struct {
	DB        interface{}
	SqlClient string
}

func NewUserDao(DB interface{}, arg ...string) *UserDao {
	userDao := &UserDao{}
	userDao.DB = DB
	if arg != nil{
		userDao.SqlClient = arg[0]
	}
	return userDao
}

// 注册用户
func (u *UserDao) CreateUser(user *DTO.User) (int64, error) {
	log.Infof("UserDao CreateUser ")
	user.CreateTime = time.Now().Unix()
	user.UpdateTime = time.Now().Unix()
	userModel := changeUserVP(user)
	client := extension.GetSQLClient(u.SqlClient)
	_, err := client.Insert(u.DB, userModel)
	if err != nil {
		log.Errorf("sql exec error info :%v", err)
		return 0, err
	}

	log.Infof(" UserDao CreateUser success")
	return userModel.Id, nil
}

// 获取用户
func (u *UserDao) GetUserById(id string) (*DTO.User, error) {
	log.Infof(" UserDao GetUserById ")
	user := &PO.User{}
	client := extension.GetSQLClient(u.SqlClient)
	_, err := client.Where(u.DB, "id = ?", id).Get(user)
	if err != nil {
		log.Errorf(" sql exec error info:%v", err)
		return nil, err
	}

	log.Infof(" UserDao GetUserById success")

	return changeUserPV(user), nil
}

// 通过用户名查询用户
func (u *UserDao) GetUserByName(name string) (*DTO.User, error) {
	log.Info("userDao GetUserByName")
	client := extension.GetSQLClient(u.SqlClient)
	user := &PO.User{}
	_, err := client.Where(u.DB, "name = ?", name).Get(user)
	if err != nil {
		log.Errorf(" sql exec error info:%v", err)
		return nil, err
	}

	log.Info("userDao GetUserByName success")
	return changeUserPV(user), nil
}

// po <-> vo
func changeUserPV(user *PO.User) *DTO.User {
	return &DTO.User{
		Id:          user.Id,
		Name:        user.Name,
		PassWord:    user.PassWord,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		CreateTime:  user.CreateTime,
		UpdateTime:  user.UpdateTime,
	}
}

func changeUserVP(user *DTO.User) *PO.User {
	return &PO.User{
		Id:          user.Id,
		Name:        user.Name,
		PassWord:    user.PassWord,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		CreateTime:  user.CreateTime,
		UpdateTime:  user.UpdateTime,
	}
}
