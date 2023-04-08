// package structure
// @Description: 非类型安全的[]string，并提供一个判断是否是StrSlice的方法
//
package structure

import "strconv"

type (
	StrSlice   []string
	StrSliceIf interface {
		IsMe()
	}
)

func (s *StrSlice) IsMe() {}

func IsStrSlice(d any) bool {
	if _, ok := any(d).(StrSliceIf); ok {
		return true
	}
	return false
}

// IntSliceToStrSlice 将标准的[]int转成[]string
func IntSliceToStrSlice(ints []int) []string {
	var strs []string
	for _, i := range ints {
		strs = append(strs, strconv.Itoa(i))
	}
	return strs
}
