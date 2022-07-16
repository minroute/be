package be

import (
	"reflect"
	"strings"
)

// 空字符串Map，主要用于辅助诊断bool
var emptyStringMap = map[string]struct{}{
	"":      {},
	"0":     {},
	"no":    {},
	"off":   {},
	"false": {},
}

// 转换bool
// 返回false："", 0, "false", "off", "no", 空slice/map,nil
// 返回true：有长度的Map/Slice/Array， 指针指向的空间不为nil
// 特例： struct 永远返回true
func beBool(v any) bool {
	if v == nil {
		return false
	}
	switch value := v.(type) {
	case bool:
		return value
	case []byte:
		str := strings.TrimSpace(string(value))
		if _, ok := emptyStringMap[strings.ToLower(str)]; ok {
			return false
		}
		return true
	case string:
		str := strings.TrimSpace(string(value))
		if _, ok := emptyStringMap[strings.ToLower(str)]; ok {
			return false
		}
		return true
	default:
		if f, ok := value.(iBool); ok {
			return f.Bool()
		}
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Ptr:
			return !rv.IsNil()
		case reflect.Map:
			fallthrough
		case reflect.Array:
			fallthrough
		case reflect.Slice:
			return rv.Len() != 0
		case reflect.Struct:
			return true
		default:
			s := strings.ToLower(String(v))
			if _, ok := emptyStringMap[s]; ok {
				return false
			}
			return true
		}
	}
}
