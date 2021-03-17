package app_config

//日志配置结构
type logStruct struct {
	LogLevel   string `ini:"level"`
	Formatter  string `ini:"formatter"`
	LogPath    string `ini:"log_path"`
	FileSuffix string `ini:"file_suffix"`
}
