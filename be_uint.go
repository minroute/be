package be

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

// beUint converts `any` to uint.
func beUint(any interface{}) uint {
	if any == nil {
		return 0
	}
	if v, ok := any.(uint); ok {
		return v
	}
	return uint(beUint64(any))
}

// beUint8 converts `any` to uint8.
func beUint8(any interface{}) uint8 {
	if any == nil {
		return 0
	}
	if v, ok := any.(uint8); ok {
		return v
	}
	return uint8(beUint64(any))
}

// beUint16 converts `any` to uint16.
func beUint16(any interface{}) uint16 {
	if any == nil {
		return 0
	}
	if v, ok := any.(uint16); ok {
		return v
	}
	return uint16(beUint64(any))
}

// beUint32 converts `any` to uint32.
func beUint32(any interface{}) uint32 {
	if any == nil {
		return 0
	}
	if v, ok := any.(uint32); ok {
		return v
	}
	return uint32(beUint64(any))
}

// beUint64 converts `any` to uint64.
func beUint64(any interface{}) uint64 {
	if any == nil {
		return 0
	}
	switch value := any.(type) {
	case int:
		return uint64(value)
	case int8:
		return uint64(value)
	case int16:
		return uint64(value)
	case int32:
		return uint64(value)
	case int64:
		return uint64(value)
	case uint:
		return uint64(value)
	case uint8:
		return uint64(value)
	case uint16:
		return uint64(value)
	case uint32:
		return uint64(value)
	case uint64:
		return value
	case float32:
		return uint64(value)
	case float64:
		return uint64(value)
	case bool:
		if value {
			return 1
		}
		return 0
	case []byte:
		//return gbinary.DecodeTobint64(value)
		var x int32
		bytesBuffer := bytes.NewBuffer(value)
		binary.Read(bytesBuffer, binary.BigEndian, &x) //大端模式
		return uint64(x)
		//@todo:return调用了unint64 。这里的代码需要进一步核对。

	default:
		if f, ok := value.(iUint64); ok {
			return f.Uint64()
		}
		s := String(value)
		// Hexadecimal
		if len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
			if v, e := strconv.ParseUint(s[2:], 16, 64); e == nil {
				return v
			}
		}
		// Octal
		if len(s) > 1 && s[0] == '0' {
			if v, e := strconv.ParseUint(s[1:], 8, 64); e == nil {
				return v
			}
		}
		// Decimal
		if v, e := strconv.ParseUint(s, 10, 64); e == nil {
			return v
		}
		// Float64
		return uint64(Float64(value))
	}
}
