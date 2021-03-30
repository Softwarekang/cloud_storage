package BO

// rpc getFileList struct
type GetFileList struct {
	UserId   int64  `json:"userId"`
	FileType string `json:"fileType"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
}
