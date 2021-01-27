package app_config

var Log logStruct

//日志配置结构
type logStruct struct {
	LogLevel  string `ini:"level"`
	Formatter string `ini:"formatter"`
	LogPath   string `ini:"log_path"`
	FileName  string `ini:"file_name"`
}
