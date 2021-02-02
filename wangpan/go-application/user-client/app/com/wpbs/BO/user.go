package BO

// gin client绑定对象
type User struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"  binding:"required"`
	PassWord    string `json:"passWord"  binding:"required"`
	PhoneNumber string `json:"phoneNumber" `
	Email       string `json:"email"`
	CreateTime  int64  `json:"createTime"`
	UpdateTime  int64  `json:"updateTime"`
}

