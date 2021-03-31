package dao

type MemoryDao struct {
	DB        interface{}
	SQLClient string
}

func NewMemoryDao(Db interface{}, arg ...string) *MemoryDao {
	memoryDao := &MemoryDao{}
	memoryDao.DB = Db
	memoryDao.SQLClient = arg[0]
	return memoryDao
}
