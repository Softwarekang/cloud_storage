package DTO

// rpc  monofile struct
type MonoFile struct {
	Id          int64
	FileName    string
	FileViewUrl string
	FileType    string
	FileSize    int64
	UserId      int64
	UserName    string
	CreateTime  int64
	UpdateTime  int64
}

func (f *MonoFile) JavaClassName() string {
	return "com.wpbs.DTO.MonoFile"
}

// rpc getFileList struct
type GetFileList struct {
	UserId   int64
	FileType string
	Page     int
	PageSize int
}

func (c *GetFileList) JavaClassName() string {
	return "com.wpbs.DTO.GetFileList"
}

type FileList struct {
	TotalCount int
	FileList   []*MonoFile
}

func (c *FileList) JavaClassName() string {
	return "com.wpbs.DTO.FileList"
}

type DeleteFile struct {
	FileIds []int64
	UserId  string
}

func (c *DeleteFile) JavaClassName() string{
	return "com.wpbs.DTO.DeleteFile"
}