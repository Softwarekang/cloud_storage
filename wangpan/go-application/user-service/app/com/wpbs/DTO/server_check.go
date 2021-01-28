package VO

// 服务探测
type ServerCheck struct {
	Code    int    // 响应码
	Message string // 响应消息
}

// 实现JavaClassName方法，同步JAVA服务
func (serverCheck *ServerCheck) JavaClassName() string{
		return "com.wpbs.DTO.ServerCheck"
}
