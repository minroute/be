package be

import (
	"database/sql"
	"reflect"
	"strconv"
)

func beNullString(v any) *sql.NullString {
	if v == nil {
		return nil
	}
	switch value := v.(type) {
	case *sql.NullString:
		return value
	case string:
		return &sql.NullString{String: value, Valid: true}
	case int64:
		return &sql.NullString{String: strconv.Itoa(beInt(value)), Valid: true}
	case int32:
		return &sql.NullString{String: strconv.Itoa(beInt(value)), Valid: true}
	case int:
		return &sql.NullString{String: strconv.Itoa(value), Valid: true}
	default:
		if f, ok := value.(iNullString); ok {
			return f.NullString()
		}
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Ptr:
			if rv.IsNil() {
				return nil
			}
			return beNullString(rv.Elem().Interface())
		case reflect.Map:
			fallthrough
		case reflect.Array:
			fallthrough
		case reflect.Slice:
			if rv.Len() == 0 {
				return nil
			}
			return beNullString(rv.Index(0).Interface())
		case reflect.Struct:
			return &sql.NullString{String: "", Valid: true}
		default:
			s := String(v)
			if s == "" {
				return nil
			}
			return &sql.NullString{String: s, Valid: true}
		}
	}
}
