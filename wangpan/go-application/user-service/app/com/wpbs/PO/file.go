package PO

// orm file struct
type file struct {
	Id          int8
	FileName    string `xorm:"varchar(30) notnull unique 'file_name'"`
	FileViewUrl string `xorm:"varchar(255) notnull  'file_view_url'"`
	FileType    string `xorm:"varchar(20) notnull 'file_type'"`
	FileSize    int8   `xorm:"'file_size'"`
	UserId      int64  `xorm:"notnull 'user_id'"`
	UserName    string `xorm:"varchar(10) notnull 'user_name'"`
	CreateTime  int64  `xorm:"not null 'create_time'"`
	UpdateTime  int64  `xorm:"not null 'create_time'"`
}

/*
CREATE TABLE `file` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `file_name` varchar(30) NOT NULL DEFAULT '' COMMENT '文件名',
  `file_view_url` varchar(255) NOT NULL DEFAULT '' COMMENT '文件浏览器URL',
  `file_type` varchar(20) NOT NULL DEFAULT '' COMMENT '文件类型',
  `file_size` int(11) NOT NULL COMMENT '文件大小(byte)',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT ' 文件所属用户ID',
  `user_name` varchar(10) NOT NULL DEFAULT '' COMMENT '文件所属用户名',
  `create_time` int(11) NOT NULL COMMENT '创建时间',
  `update_time` int(11) NOT NULL COMMENT '更新事件',
  PRIMARY KEY (`id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
