package common

import "reflect"

// 检测是否为空
func IsNilString(str interface{}) (bool, error) {
	if reflect.TypeOf(str).Kind() != reflect.String {
		return false, new(Errors).New("interface type must be string")
	}

	if str == nil || str == "" {
		return true, nil
	}

	return false, nil
}