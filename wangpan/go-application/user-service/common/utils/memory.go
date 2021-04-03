package utils

func MbToByte(mb int64) int64 {
	if mb > 9223372036854775807/(1024*1024*1024) {
		// todo log补充
		return 0
	}
	
	return mb * 1024 * 1024
}
