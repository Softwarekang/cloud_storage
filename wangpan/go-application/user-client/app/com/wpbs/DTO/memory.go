package DTO

 type Memory struct {
	Id             int64
	UserName       string
	UserId         int64
	ConsumeMemory  int64
	MemoryCapacity int64
	CreateTime     int64
	UpdateTime     int64
}

func (f *Memory) JavaClassName() string {
	return "com.wpbs.DTO.Memory"
}