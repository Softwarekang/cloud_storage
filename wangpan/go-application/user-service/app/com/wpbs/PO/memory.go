package PO

// 内存表 orm
type Memory struct {
	Id             int64
	UserName       string `xorm:"varchar(10) notnull 'user_name'"`
	UserId         int64  `xorm:"notnull 'user_id'"`
	ConsumeMemory  int64  `xorm:"notnull 'consume_memory'"`
	MemoryCapacity int64  `xorm:"notnull 'memory_capacity'"`
	CreateTime     int64  `xorm:"not null 'create_time'"`
	UpdateTime     int64  `xorm:"not null 'update_time'"`
}

/*
CREATE TABLE `memory` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_name` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '鐢ㄦ埛濮撳悕',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
  `consume_memory` bigint(20) NOT NULL COMMENT '占用内存(byte)',
  `memory_capacity` bigint(20) unsigned NOT NULL COMMENT '内存容量(byte)',
  `create_time` int(11) NOT NULL COMMENT '创建时间',
  `update_time` int(11) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8;
*/
