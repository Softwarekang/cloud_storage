package dao

import (
	"time"
	"user-service/app/com/wpbs/DTO"
	"user-service/app/com/wpbs/PO"
	"user-service/app/com/wpbs/config/database"
	"user-service/common"
)

var (
	log = common.GetLogger()
	DB  = database.GetDB()
)

type UserDao struct {
}

// 注册用户
func (u *UserDao) CreateUser(user *DTO.User) (int64, error) {
	log.Infof("UserDao CreateUser ")
	user.CreateTime = time.Now().Unix()
	user.UpdateTime = time.Now().Unix()
	userModel := changeUserVP(user)
	_, err := DB.Insert(userModel)
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
	_, err := DB.Where("id = ?", id).Get(user)
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
	user := &PO.User{}
	_, err := DB.Where("name = ?", name).Get(user)
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
