package be

//// Bytes converts and returns `v` as []byte.
//func Bytes(v any) []byte {
//
//}

// String converts and returns `v` as string.
func String(v any) string {
	return beString(v)
}

//// Bool converts and returns `v` as bool.
func Bool(v any) bool {
	return beBool(v)
}

//
//// Int converts and returns `v` as int.
//func Int(v any) int {
//}
//
//// Int8 converts and returns `v` as int8.
//func Int8(v any) int8 {
//}
//
//// Int16 converts and returns `v` as int16.
//func Int16(v any) int16 {
//}
//
//// Int32 converts and returns `v` as int32.
//func Int32(v any) int32 {
//}
//
//// Int64 converts and returns `v` as int64.
//func Int64(v any) int64 {
//}
//
//// Uint converts and returns `v` as uint.
//func Uint(v any) uint {
//}
//
//// Uint8 converts and returns `v` as uint8.
//func Uint8(v any) uint8 {
//}
//
//// Uint16 converts and returns `v` as uint16.
//func Uint16(v any) uint16 {
//}
//
//// Uint32 converts and returns `v` as uint32.
//func Uint32(v any) uint32 {
//}
//
//// Uint64 converts and returns `v` as uint64.
//func Uint64(v any) uint64 {
//}
//
//// Float32 converts and returns `v` as float32.
//func Float32(v any) float32 {
//}
//
//// Float64 converts and returns `v` as float64.
//func Float64(v any) float64 {
//}
