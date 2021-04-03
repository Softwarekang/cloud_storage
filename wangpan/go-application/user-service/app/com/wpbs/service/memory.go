package service

import (
	"context"
	"user-service/app/com/wpbs/DTO"
	"user-service/app/com/wpbs/store"
)

type MemoryService struct {
}

func (m *MemoryService) GetMemory(ctx context.Context, userId int64) (*DTO.Memory, error) {
	log.Infof("MemoryService GetMemory req: userId:%v", userId)
	engine := store.DBClient.Begin()
	memory, err := store.DBClient.Memory(engine).GetMemoryByUserId(userId)
	if err != nil {
		return nil, err
	}

	log.Infof("MemoryService GetMemory rsp model:%v", *memory)
	return memory, nil
}

func (m *MemoryService) Reference() string {
	return "MemoryService"
}
