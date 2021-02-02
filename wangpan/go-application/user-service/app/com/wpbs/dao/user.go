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
	log.Infof("traceId:%v UserDao CreateUser ", user.TraceId)
	user.CreateTime = time.Now().Unix()
	user.UpdateTime = time.Now().Unix()
	userModel := changeUserVP(user)
	_, err := DB.Insert(userModel)
	if err != nil {
		log.Errorf("traceId:%v sql exec error info :", user.TraceId, err)
		return 0, err
	}

	log.Infof("traceId:%v UserDao CreateUser success", user.TraceId)
	return userModel.Id, nil
}

// 获取用户
func (u *UserDao) GetUserById(id int64, traceId string) (*DTO.User, error) {
	log.Infof("traceId:%v UserDao GetUserById ", traceId)
	user := &PO.User{}
	_, err := DB.Where("id = ?", id).Get(user)
	if err != nil {
		log.Errorf("traceId:%v sql exec error info:", traceId, err)
		return nil, err
	}

	log.Infof("traceId:%v UserDao GetUserById success", traceId)

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
