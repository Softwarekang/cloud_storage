package dao

import (
	"fmt"
	"strings"
)

func WithSQLParam(sql, paramName string, paramValue interface{}) string {
	if paramValue == nil || paramValue == ""{
		return sql
	}

	if strings.Contains(sql, "where") {
		sql = fmt.Sprintf("%s and %s = '%v'", sql, paramName, paramValue)
	} else {
		sql = fmt.Sprintf("%s where %s = '%v'", sql, paramName, paramValue)
	}

	return sql
}

func IDToStr(ids []int64) string {
	strs := "("
	for index, value := range ids {
		if index == len(ids)-1 {
			strs += fmt.Sprintf("%d)", value)
			break
		}

		strs += fmt.Sprintf("%d,", value)
	}

	return strs
}
