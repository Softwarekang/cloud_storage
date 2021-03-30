package utils

import (
	"reflect"
	"strconv"
	"user-client/common/errors"
)

// 检测是否为空
func IsNilString(str interface{}) (bool, error) {
	if reflect.TypeOf(str).Kind() != reflect.String {
		return false, new(errors.Errors).New("interface type must be string")
	}

	if str == nil || str == "" {
		return true, nil
	}

	return false, nil
}

// string[] -> int64[]
func StringArrayToInt64Array(ids []string, rsp *[]int64) error {
	if ids == nil || len(ids) == 0 {
		return new(errors.Errors).New("ids must not be empty or nil")
	}

	*rsp = make([]int64, 0, len(ids))
	for _, v := range ids {
		intV, _ := strconv.ParseInt(v, 10, 64)
		*rsp = append(*rsp, intV)
	}

	return nil
}

