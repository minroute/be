// package structure
// @Description: 非类型安全的[]string，并提供一个判断是否是StrSlice的方法
//
package structure

import (
	"unsafe"
)

type (
	IntSlice   []int
	IntSliceIf interface {
		IsMe()
	}
)

func (s *IntSlice) IsMe() {}

func IsIntSlice(d any) bool {
	// if kit.IsPtr(d) {
	// 	panic("IsIntSlice() only accept non-pointer type")
	// }
	if _, ok := any(d).(IntSliceIf); ok {
		return true
	}
	return false
}

// ToIntSlice 将标准的[]int转成*IntSlice ,值传递
func ToIntSlice(s []int) *IntSlice {
	return (*IntSlice)(unsafe.Pointer(&s)) // return IntSlice(s)
}

// ToIntSliceFromPtrAddr 将标准的[]int转成*IntSlice,  指针传递
func ToIntSliceFromPtrAddr(s *[]int) *IntSlice {
	return (*IntSlice)(unsafe.Pointer(s))
}
