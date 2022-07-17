package be

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

// Int converts `v` to int.
func beInt(av any) int {
	if av == nil {
		return 0
	}
	if v, ok := av.(int); ok {
		return v
	}
	return int(beInt64(av))
}

// Int8 converts `v` to int8.
func beInt8(av any) int8 {
	if av == nil {
		return 0
	}
	if v, ok := av.(int8); ok {
		return v
	}
	return int8(beInt64(av))
}

// Int16 converts `v` to int16.
func beInt16(av any) int16 {
	if av == nil {
		return 0
	}
	if v, ok := av.(int16); ok {
		return v
	}
	return int16(beInt64(av))
}

// Int32 converts `v` to int32.
func beInt32(av any) int32 {
	if av == nil {
		return 0
	}
	if v, ok := av.(int32); ok {
		return v
	}
	return int32(beInt64(av))
}

// Int64 converts `v` to int64.
func beInt64(v any) int64 {
	if v == nil {
		return 0
	}
	switch value := v.(type) {
	case int:
		return int64(value)
	case int8:
		return int64(value)
	case int16:
		return int64(value)
	case int32:
		return int64(value)
	case int64:
		return value
	case uint:
		return int64(value)
	case uint8:
		return int64(value)
	case uint16:
		return int64(value)
	case uint32:
		return int64(value)
	case uint64:
		return int64(value)
	case float32:
		return int64(value)
	case float64:
		return int64(value)
	case bool:
		if value {
			return 1
		}
		return 0

	// byte数组转int64，
	// 知识点：byte其实是uint8的别名，byte 和 uint8 之间可以直接进行互转。
	//        目前来只能将0~255范围的int转成byte。因为超出这个范围，go在转换的时候，就会把多出来数据扔掉;
	//        如果需要将int32转成byte类型，我们只需要一个长度为4的[]byte数组
	// 参考：  https://studygolang.com/articles/16154
	case []byte:
		var x int32
		bytesBuffer := bytes.NewBuffer(value)
		binary.Read(bytesBuffer, binary.BigEndian, &x) //大端模式
		return int64(x)

	// @todo：map、slice、struct转int64，此处并为处理
	default:
		if f, ok := value.(iInt64); ok {
			return f.Int64()
		}
		s := String(value)
		isMinus := false
		if len(s) > 0 {
			if s[0] == '-' {
				isMinus = true
				s = s[1:]
			} else if s[0] == '+' {
				s = s[1:]
			}
		}
		// Hexadecimal 16进制
		if len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
			if v, e := strconv.ParseInt(s[2:], 16, 64); e == nil {
				if isMinus {
					return -v
				}
				return v
			}
		}
		// Octal 8进制
		if len(s) > 1 && s[0] == '0' {
			if v, e := strconv.ParseInt(s[1:], 8, 64); e == nil {
				if isMinus {
					return -v
				}
				return v
			}
		}
		// Decimal 10进制
		if v, e := strconv.ParseInt(s, 10, 64); e == nil {
			if isMinus {
				return -v
			}
			return v
		}
		// Float64
		return int64(Float64(value))
	}
}
