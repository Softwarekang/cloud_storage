package PO

// 数据库对应结构
type User struct {
	Id           int64
	Name         string `xorm:"varchar(10) notnull unique 'name'"`
	PassWord     string `xorm:"varchar(20) not null 'password'"`
	PhoneNumber  string `xorm:"varchar(15) 'phone_number'"`
	Email        string `xorm:"varchar(20) 'email'"`
	HeadImageUrl string `xorm:"varchar(50) notnull 'head_image_url'"`
	CreateTime   int64  `xorm:"notnull 'create_time'"`
	UpdateTime   int64  `xorm:"notnull 'update_time'"`
}

/*
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '全局ID',
  `name` varchar(10) NOT NULL COMMENT '用户姓名',
  `password` varchar(20) NOT NULL COMMENT '用户密码',
  `phone_number` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '鐢ㄦ埛鐢佃瘽',
  `email` varchar(25) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '鐢ㄦ埛閭鍦板潃',
  `head_image_url` varchar(50) NOT NULL DEFAULT '' COMMENT '用户头像路径',
  `create_time` int(11) NOT NULL COMMENT '创建时间',
  `update_time` int(11) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=216 DEFAULT CHARSET=utf8;
*/
