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
	SQLClient string
}

func NewUserDao(DB interface{}, arg ...string) *UserDao {
	userDao := &UserDao{}
	userDao.DB = DB
	if arg != nil {
		userDao.SQLClient = arg[0]
	}
	return userDao
}

// 注册用户
func (u *UserDao) CreateUser(user *DTO.User) (int64, error) {
	log.Infof("UserDao CreateUser model:%v", *user)
	user.CreateTime = time.Now().Unix()
	user.UpdateTime = time.Now().Unix()
	userModel := changeUserVP(user)
	client := extension.GetSQLClient(u.SQLClient)
	_, err := client.Insert(u.DB, userModel)
	if err != nil {
		log.Errorf("sql exec error info :%v,model:%v", err, *userModel)
		return 0, err
	}

	log.Infof(" UserDao CreateUser success, , model:%v", *userModel)
	return userModel.Id, nil
}

// 获取用户
func (u *UserDao) GetUserById(id string) (*DTO.User, error) {
	log.Infof(" UserDao GetUserById id:%v ", id)
	user := &PO.User{}
	client := extension.GetSQLClient(u.SQLClient)
	_, err := client.Where(u.DB, "id = ?", id).Get(user)
	if err != nil {
		log.Errorf(" sql exec error info:%v, id:%v", err, id)
		return nil, err
	}

	log.Infof(" UserDao GetUserById success rsp model:%v", user)

	return changeUserPV(user), nil
}

// 通过用户名查询用户
func (u *UserDao) GetUserByName(name string) (*DTO.User, error) {
	log.Info("userDao GetUserByName name:%v", name)
	client := extension.GetSQLClient(u.SQLClient)
	user := &PO.User{}
	_, err := client.Where(u.DB, "name = ?", name).Get(user)
	if err != nil {
		log.Errorf(" sql exec error info:%v, name:%v", err, name)
		return nil, err
	}

	log.Info("userDao GetUserByName success rsp model:%v", user)
	return changeUserPV(user), nil
}

// 更新用户信息
func (u *UserDao) UpdateUser(user *DTO.User) error {
	log.Info("userDao UpdateUser param:%v", user)
	userModel := changeUserVP(user)
	client := extension.GetSQLClient(u.SQLClient)
	if _, err := client.ID(u.DB, user.Id).Update(userModel); err != nil {
		log.Errorf("UserDao UpdateUser error:%v, model:%v", err, userModel)
		return err
	}

	log.Infof("UserDao UpdateUser success")
	return nil
}

// po <-> vo
func changeUserPV(user *PO.User) *DTO.User {
	return &DTO.User{
		Id:           user.Id,
		Name:         user.Name,
		PassWord:     user.PassWord,
		PhoneNumber:  user.PhoneNumber,
		Email:        user.Email,
		HeadImageUrl: user.HeadImageUrl,
		CreateTime:   user.CreateTime,
		UpdateTime:   user.UpdateTime,
	}
}

func changeUserVP(user *DTO.User) *PO.User {
	return &PO.User{
		Id:           user.Id,
		Name:         user.Name,
		PassWord:     user.PassWord,
		PhoneNumber:  user.PhoneNumber,
		Email:        user.Email,
		HeadImageUrl: user.HeadImageUrl,
		CreateTime:   user.CreateTime,
		UpdateTime:   user.UpdateTime,
	}
}
