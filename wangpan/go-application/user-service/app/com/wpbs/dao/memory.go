package dao

import (
	"time"
	"user-service/app/com/wpbs/DTO"
	"user-service/app/com/wpbs/PO"
	"user-service/common/extension"
)

type MemoryDao struct {
	DB        interface{}
	SQLClient string
}

func NewMemoryDao(Db interface{}, arg ...string) *MemoryDao {
	memoryDao := &MemoryDao{}
	memoryDao.DB = Db
	if arg != nil {
		memoryDao.SQLClient = arg[0]
	}
	return memoryDao
}

// 创建内存信息
func (m *MemoryDao) CreateMemory(helper *DTO.Memory) error {
	log.Infof("MemoryDao CreateMemory model:%v", *helper)
	helper.CreateTime = time.Now().Unix()
	helper.UpdateTime = time.Now().Unix()
	memoryModel := ChangeMemoryDP(helper)
	_, err := extension.GetSQLClient(m.SQLClient).Insert(m.DB, memoryModel)
	if err != nil {
		log.Errorf("MemoryDao CreateMemory error:%v, model:%v", err, memoryModel)
		return err
	}

	log.Infof("MemoryDao CreateMemory success model:%v", *memoryModel)
	return nil
}

// 更新内存状态
func (m *MemoryDao) UpdateMemory(userId, consumeMemory int64) error {
	log.Infof("MemoryDao UpdateMemory param: userId:%v,consumeMemory:%v", userId, consumeMemory)
	nowTime := time.Now().Unix()
	sql := WithSQLParam("update memory set consume_memory =  ?, update_time = ?", "user_id", userId)
	if _, err := extension.GetSQLClient(m.SQLClient).SQL(m.DB, sql, consumeMemory, nowTime).Execute(); err != nil {
		log.Errorf("MemoryDao UpdateMemory error:%v, sql:%v", err, sql)
		return err
	}

	log.Infof("MemoryDao UpdateMemory success, sql:%v", sql)
	return nil
}

// 查询当前容量
func (m *MemoryDao) GetMemoryByUserId(userId int64) (*DTO.Memory, error) {
	log.Infof("MemoryDao GetConsumeMemoryByUserId param:userId:%v", userId)
	sql := WithSQLParam("select * from memory", "user_id", userId)
	memory := &PO.Memory{}
	if _, err := extension.GetSQLClient(m.SQLClient).SQL(m.DB, sql).Get(memory); err != nil {
		log.Errorf("MemoryDao GetConsumeMemoryByUserId error:%v,sql:%v", err, sql)
		return nil, err
	}

	log.Infof("MemoryDao GetConsumeMemoryByUserId success, sql: %v,rsp:%v", sql, *memory)
	return ChangeMemoryPD(memory), nil
}

func ChangeMemoryDP(helper *DTO.Memory) *PO.Memory {
	return &PO.Memory{
		Id:             helper.Id,
		UserName:       helper.UserName,
		UserId:         helper.UserId,
		ConsumeMemory:  helper.ConsumeMemory,
		MemoryCapacity: helper.MemoryCapacity,
		CreateTime:     helper.CreateTime,
		UpdateTime:     helper.UpdateTime,
	}
}

func ChangeMemoryPD(memory *PO.Memory) *DTO.Memory {
	return &DTO.Memory{
		Id:             memory.Id,
		UserName:       memory.UserName,
		UserId:         memory.UserId,
		ConsumeMemory:  memory.ConsumeMemory,
		MemoryCapacity: memory.MemoryCapacity,
		CreateTime:     memory.CreateTime,
		UpdateTime:     memory.UpdateTime,
	}
}
