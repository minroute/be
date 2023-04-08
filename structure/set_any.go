// package structure
// @Description: 可以添加有限类型的Set接口，支持并发安全
// 只有支持comparable的类型才可以顺利添加进来。
// comparable的类型有：boolean, numeric, string, pointer, channel, interface, 以及仅包含这些类型的struct和array 。
// 使用方法看container_any_set_test.go即可
package structure

import (
	"encoding/json"
	"fmt"
	"sync"
)

type AnySetCore struct {
	Data   map[any]struct{}
	rwlock *sync.RWMutex
}

// AnySet 实例化任意类型的Set
func AnySet(safe ...bool) *AnySetCore {
	isSafe := false
	if len(safe) > 0 && safe[0] {
		isSafe = safe[0]
	}
	if isSafe {
		return &AnySetCore{
			Data:   make(map[any]struct{}),
			rwlock: new(sync.RWMutex), // 通过new 来分配类型的指针。new(T) 返回一个指向T类型的指针，该指针指向的值为T类型的零值。这里用来优化性能
		}
	} else {
		return &AnySetCore{
			Data:   make(map[any]struct{}),
			rwlock: nil,
		}
	}
}

// Add 添加1-n个元素到集合，元素可以是任意类型
// map key 必须是可比较的类型，语言规范中定义了可比较的类型：boolean, numeric, string, pointer, channel, interface, 以及仅包含这些类型的struct和array 。
// 不能作为map key的类型有：slice，map, function。
// As mentioned earlier, map keys may be of any type that is comparable.
// The language spec defines this precisely, but in short, comparable types are
// boolean, numeric, string, pointer, channel, and interface types, and structs or arrays that contain only those types.
// Notably absent from the list are slices, maps, and functions; these types cannot be compared using ==, and may not be used as map keys.
func (s *AnySetCore) Add(items ...any) *AnySetCore {
	if s.rwlock != nil {
		s.rwlock.Lock()
		defer s.rwlock.Unlock()
	}
	for _, item := range items {
		s.Data[item] = struct{}{}
	}

	return s
}

// Delete  从集合中删除1-n个元素，元素可以是任意类型
func (s *AnySetCore) Delete(items ...any) *AnySetCore {
	if s.rwlock != nil {
		s.rwlock.Lock()
		defer s.rwlock.Unlock()
	}
	for _, item := range items {
		delete(s.Data, item)
	}
	return s
}

// Clear 清空集合
func (s *AnySetCore) Clear() *AnySetCore {
	if s.rwlock != nil {
		s.rwlock.Lock()
		defer s.rwlock.Unlock()
	}
	for k := range s.Data {
		delete(s.Data, k)
	}
	return s
}

// Iterate  遍历集合，无返回
func (s *AnySetCore) Iterate(it func(any)) {
	if s.rwlock != nil {
		s.rwlock.RLock()
		defer s.rwlock.RUnlock()
	}
	for k := range s.Data {
		it(k)
	}

}

// Len 返回集合的元素个数
func (s *AnySetCore) Len() int {
	if s.rwlock != nil {
		s.rwlock.RLock()
		defer s.rwlock.RUnlock()
	}
	return len(s.Data)
}

// Has 判断集合中是否存在某个元素
func (s *AnySetCore) Has(item any) bool {
	if s.rwlock != nil {
		s.rwlock.RLock()
		defer s.rwlock.RUnlock()
	}
	_, ok := s.Data[item]
	return ok
}

// String 返回集合的字符串表示,实现了fmt.Stringer接口
func (s *AnySetCore) String() string {
	return fmt.Sprint(s.ToSlice())
}

// ToSlice 返回集合的切片
func (s *AnySetCore) ToSlice() []any {
	if s.rwlock != nil {
		s.rwlock.RLock()
		defer s.rwlock.RUnlock()
	}
	ret := make([]any, len(s.Data))
	i := 0
	for k := range s.Data {
		ret[i] = k
		i++
	}
	return ret
}

// MarshalJSON 返回json字节切片和错误信息，实现了json.Marshaler接口
func (s *AnySetCore) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.ToSlice())
}

// JsonString 返回json字符串,如果出错，则返回空字符串
func (s *AnySetCore) JsonString() string {
	r, e := s.MarshalJSON()
	if e != nil {
		return ""
	}
	return string(r)
}
