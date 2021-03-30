package utils

import "path"

// 获取文件后缀
func GetSuffix(fileName string) string {
	suffix := path.Ext(fileName)
	if suffix == "" {
		suffix = ".default"
	}
	return suffix[1:]
}

// 根据后缀获取文件类型
func GetTypeBySuffix(suffix string) string {
	typeMap := make(map[string]string)
	// 特殊文本类型
	typeMap["html"] = "text"
	typeMap["xml"] = "text"
	typeMap["json"] = "text"
	typeMap["pdf"] = "text"
	typeMap["doc"] = "text"
	typeMap["txt"] = "text"
	// 图片类型
	typeMap["gif"] = "image"
	typeMap["jpeg"] = "image"
	typeMap["png"] = "image"
	typeMap["jpg"] = "image"
	// 音乐
	typeMap["mp3"] = "music"
	// 视频
	typeMap["mp4"] = "video"
	// 其他
	typeMap["default"] = "others"
	if res, ok := typeMap[suffix]; ok {
		return res
	}
	return typeMap["default"]
}
