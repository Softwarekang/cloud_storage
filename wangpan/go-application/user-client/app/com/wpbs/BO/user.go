package BO

// gin client绑定对象
type User struct {
	Id           int64  `json:"id"`
	Name         string `form:"userName" json:"userName"  binding:"required"`
	PassWord     string `form:"passWord" json:"passWord"  binding:"required"`
	PhoneNumber  string `json:"phoneNumber" `
	Email        string `json:"email"`
	HeadImageUrl string `json:"headImageUrl"`
	CreateTime   int64  `json:"createTime"`
	UpdateTime   int64  `json:"updateTime"`
}
