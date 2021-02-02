package PO

// 数据库对应结构
type User struct {
	Id          int64
	Name        string `xorm:"varchar(10) notnull unique 'name'`
	PassWord    string `xorm:"varchar(20) not null 'password'"`
	PhoneNumber string `xorm:"varchar(15) 'phone_number'"`
	Email       string `xorm:"varchar(20) 'email'"`
	CreateTime  int64    `xorm:"notnull 'create_time'"`
	UpdateTime  int64    `xorm:"notnull 'update_time'"`
}
