package service

import (
	"context"
	"user-client/app/com/wpbs/DTO"
)

type MemoryService struct {
	GetMemory func(ctx context.Context, userId int64, rsp *DTO.Memory) error
}

func (m *MemoryService) Reference() string {
	return "MemoryService"
}
