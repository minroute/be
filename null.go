// Package be
// @Description: 以下是null系列的转换方法，用于将某个变量转换为对应的null类型，方便数据库操作
//
package be

import "database/sql"

func NullString(v any) *sql.NullString {
	return beNullString(v)
}

func NullBool(v any) *sql.NullBool {
	return beNullBool(v)
}

func NullInt64(v any) *sql.NullInt64 {
	return beNullInt64(v)
}

func NullInt32(v any) *sql.NullInt32 {
	return beNullInt32(v)
}

func NullInt16(v any) *sql.NullInt16 {
	return beNullInt16(v)
}

func NullTime(v any) *sql.NullTime {
	return beNullTime(v)
}
