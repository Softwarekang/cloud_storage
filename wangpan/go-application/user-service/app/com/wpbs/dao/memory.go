package dao

import (
	"user-service/app/com/wpbs/PO"
	"user-service/app/com/wpbs/dao/helper"
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
func (m *MemoryDao) CreateMemory(helper *helper.CreateMemoryHelper) error {
	log.Infof("MemoryDao CreateMemory model:%v", *helper)
	memoryModel := ChangeMemoryHP(helper)
	_, err := extension.GetSQLClient(m.SQLClient).Insert(m.DB, memoryModel)
	if err != nil {
		log.Errorf("MemoryDao insert error:%v, model:%v", err, memoryModel)
		return err
	}

	log.Infof("MemoryDao CreateMemory success model:%v", *memoryModel)
	return nil
}

func ChangeMemoryHP(helper *helper.CreateMemoryHelper) *PO.Memory {
	return &PO.Memory{
		Id:             helper.Id,
		UserName:       helper.UserName,
		UserId:         helper.UserId,
		MemoryCapacity: helper.MemoryCapacity,
		CreateTime:     helper.CreateTime,
		UpdateTime:     helper.UpdateTime,
	}
}

func ChangeMemoryPH(memory *PO.Memory) *helper.CreateMemoryHelper {
	return &helper.CreateMemoryHelper{
		Id:             memory.Id,
		UserName:       memory.UserName,
		UserId:         memory.UserId,
		MemoryCapacity: memory.MemoryCapacity,
		CreateTime:     memory.CreateTime,
		UpdateTime:     memory.UpdateTime,
	}
}
