package be

import (
	"encoding/binary"
	"math"
	"strconv"
)

// Float32 converts `any` to float32.
func beFloat32(any interface{}) float32 {
	if any == nil {
		return 0
	}
	switch value := any.(type) {
	case float32:
		return value
	case float64:
		return float32(value)
	case []byte:
		// 注意此处，是小端序，且leFillUpSize方法是辅助用途
		return math.Float32frombits(binary.LittleEndian.Uint32(leFillUpSize(value, 4)))
	default:
		if f, ok := value.(iFloat32); ok {
			return f.Float32()
		}
		v, _ := strconv.ParseFloat(String(any), 64)
		return float32(v)
	}
}

// Float64 converts `any` to float64.
func beFloat64(any interface{}) float64 {
	if any == nil {
		return 0
	}
	switch value := any.(type) {
	case float32:
		return float64(value)
	case float64:
		return value
	case []byte:
		// 注意此处，是小端序，且leFillUpSize方法是辅助用途
		return math.Float64frombits(binary.LittleEndian.Uint64(leFillUpSize(value, 8)))
	default:
		if f, ok := value.(iFloat64); ok {
			return f.Float64()
		}
		v, _ := strconv.ParseFloat(String(any), 64)
		return v
	}
}
