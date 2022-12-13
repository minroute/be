package be

import (
	"database/sql"
	"reflect"
	"strconv"
	"strings"
)

/**
 * beNullBool：将任意类型转换为sql.NullBool。转换原则；非0即真，非空即真，非nil/null即真。
 * @param v any：任意类型
 * @return *sql.NullBool：sql.NullBool
 * 基本类型遵从非0即真，其他类型遵从非空即真活是否实现自定义接口iNullBool
 */
func beNullBool(v any) *sql.NullBool {
	if v == nil {
		return &sql.NullBool{Bool: false, Valid: true}
	}
	switch value := v.(type) {
	case *sql.NullBool:
		return value
	case string:
		if value == "" || strings.ToUpper(value) == "NULL" || strings.ToUpper(value) == "NIL" || value == "0" {
			return &sql.NullBool{Bool: false, Valid: true}
		}
		return &sql.NullBool{Bool: true, Valid: true}
	case int64:
	case int32:
	case int:
		if value > 0 {
			return &sql.NullBool{Bool: true, Valid: true}
		}
		return &sql.NullBool{Bool: false, Valid: true}
	default:
		if f, ok := value.(iNullBool); ok {
			return f.NullBool()
		}
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Ptr:
			// 空指针为false，非空指针为true
			if rv.IsNil() {
				return &sql.NullBool{Bool: false, Valid: true}
			}
			return &sql.NullBool{Bool: true, Valid: true}
		case reflect.Map:
			fallthrough
		case reflect.Array:
			fallthrough
		case reflect.Slice:
			// 非空为true，空为false
			if rv.Len() == 0 {
				return &sql.NullBool{Bool: false, Valid: true}
			}
			return &sql.NullBool{Bool: true, Valid: true}
		case reflect.Struct:
			// 空结构返回false，非空结构返回true
			if rv.NumField() == 0 {
				return &sql.NullBool{Bool: false, Valid: true}
			}
			return &sql.NullBool{Bool: true, Valid: true}
		default:
			// 强制转字符串后，如果为空字符串，则返回false，否则调用strconv.ParseBool进行转换。
			// 如转换失败，则返回false
			s := String(v)
			if s == "" {
				return &sql.NullBool{Bool: false, Valid: true}
			}
			if b, e := strconv.ParseBool(s); e == nil {
				return &sql.NullBool{Bool: b, Valid: true}
			}
			return &sql.NullBool{Bool: false, Valid: true}
		}
	}
	return &sql.NullBool{Bool: false, Valid: true}
}
