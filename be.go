package be

// Bytes converts and returns `v` as []byte.
func Bytes(v any) []byte {
	return beBytes(v)
}

// String converts and returns `v` as string.
func String(v any) string {
	return beString(v)
}

// Bool converts and returns `v` as bool.
func Bool(v any) bool {
	return beBool(v)
}

// Int converts and returns `v` as int.
func Int(v any) int {
	return beInt(v)
}

// Int8 converts and returns `v` as int8.
func Int8(v any) int8 {
	return beInt8(v)
}

// Int16 converts and returns `v` as int16.
func Int16(v any) int16 {
	return beInt16(v)
}

// Int32 converts and returns `v` as int32.
func Int32(v any) int32 {
	return beInt32(v)
}

// Int64 converts and returns `v` as int64.
func Int64(v any) int64 {
	return beInt64(v)
}

// Uint converts and returns `v` as uint.
func Uint(v any) uint {
	return beUint(v)
}

// Uint8 converts and returns `v` as uint8.
func Uint8(v any) uint8 {
	return beUint8(v)
}

// Uint16 converts and returns `v` as uint16.
func Uint16(v any) uint16 {
	return beUint16(v)
}

// Uint32 converts and returns `v` as uint32.
func Uint32(v any) uint32 {
	return beUint32(v)
}

// Uint64 converts and returns `v` as uint64.
func Uint64(v any) uint64 {
	return beUint64(v)
}

// Float32 converts and returns `v` as float32.
func Float32(v any) float32 {
	return beFloat32(v)
}

// Float64 converts and returns `v` as float64.
func Float64(v any) float64 {
	return beFloat64(v)
}
