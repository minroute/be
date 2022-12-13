// 定义了一批接口，用来辅助做类型断言所用。
// 如某个变量实现了istring接口，那么beString方法将调用istring.String方法来获取字符串。

package be

import "database/sql"

//
// 以下定义一批常规类型接口，用来检测某个变量是否实现了某个接口。方便转换的时候做客制化
// 例子：如某个变量实现了 iString 接口，那么beString方法将调用 iString.String方法来获取字符串。
//
//

// iString is used for type assert api for String().
type iString interface {
	String() string
}

// iBool is used for type assert api for Bool().
type iBool interface {
	Bool() bool
}

// iInt64 is used for type assert api for Int64().
type iInt64 interface {
	Int64() int64
}

// iUint64 is used for type assert api for Uint64().
type iUint64 interface {
	Uint64() uint64
}

// iFloat32 is used for type assert api for Float32().
type iFloat32 interface {
	Float32() float32
}

// iFloat64 is used for type assert api for Float64().
type iFloat64 interface {
	Float64() float64
}

// iError is used for type assert api for Error().
type iError interface {
	Error() string
}

// iBytes is used for type assert api for Bytes().
type iBytes interface {
	Bytes() []byte
}

// iInterface is used for type assert api for Interface().
type iInterface interface {
	Interface() interface{}
}

// iInterfaces is used for type assert api for Interfaces().
type iInterfaces interface {
	Interfaces() []interface{}
}

// iFloats is used for type assert api for Floats().
type iFloats interface {
	Floats() []float64
}

// iInts is used for type assert api for Ints().
type iInts interface {
	Ints() []int
}

// iStrings is used for type assert api for Strings().
type iStrings interface {
	Strings() []string
}

// iUints is used for type assert api for Uints().
type iUints interface {
	Uints() []uint
}

// iMapStrAny is the interface support for converting struct parameter to map.
type iMapStrAny interface {
	MapStrAny() map[string]interface{}
}

// iUnmarshalValue is the interface for custom defined types customizing value assignment.
// Note that only pointer can implement interface iUnmarshalValue.
type iUnmarshalValue interface {
	UnmarshalValue(interface{}) error
}

// iUnmarshalText is the interface for custom defined types customizing value assignment.
// Note that only pointer can implement interface iUnmarshalText.
type iUnmarshalText interface {
	UnmarshalText(text []byte) error
}

// iUnmarshalText is the interface for custom defined types customizing value assignment.
// Note that only pointer can implement interface iUnmarshalJSON.
type iUnmarshalJSON interface {
	UnmarshalJSON(b []byte) error
}

// iSet is the interface for custom value assignment.
type iSet interface {
	Set(value interface{}) (old interface{})
}

//
// 以下定义一批Null类型的接口，用来检测某个变量是否实现了某个接口。方便转换的时候做客制化
// 如某个变量实现了 iNullString 接口，那么beNullString方法将调用 iNullString.NullString方法来获取字符串。
//
//
// iNullString is the interface for custom value assignment.
type iNullString interface {
	NullString() *sql.NullString
}

// iNullBool is the interface for custom value assignment.
type iNullBool interface {
	NullBool() *sql.NullBool
}

// iNullInt64 is the interface for custom value assignment.
type iNullInt64 interface {
	NullInt64() *sql.NullInt64
}

type iNullInt32 interface {
	NullInt32() *sql.NullInt32
}

type iNullInt16 interface {
	NullInt16() *sql.NullInt16
}

type iNullFloat64 interface {
	NullFloat64() *sql.NullFloat64
}

type iNullTime interface {
	NullTime() *sql.NullTime
}

type iNullByte interface {
	NullByte() *sql.NullByte
}
