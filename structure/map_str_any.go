// package structure
// @Description: 非类型安全的map[string]interface{}，并提供一个判断是否是StrAnyMap的方法
//
package structure

type (
	StrAnyMap   map[string]any
	StrAnyMapIf interface {
		IsMe()
	}
)

func (s *StrAnyMap) IsMe() {}

// IsStrAnyMap 判断是否是StrAnyMap,
// @usage: IsStrAnyMap(StrAnyMap{"a": 1})
// @usage: IsStrAnyMap(&StrAnyMap{"a": 1})
func IsStrAnyMap(d any) bool {
	if IsPtr(d) {
		if _, ok := any(d).(*StrAnyMap); ok {
			return true
		}
	} else {
		if _, ok := any(d).(StrAnyMap); ok {
			return true
		}
	}
	return false
}
