package structure

import "fmt"

type Set[T comparable] map[T]struct{}

// Add 添加1-n个元素到集合，元素可以是任意类型
func (s Set[T]) Add(values ...T) Set[T] {
	for i := 0; i < len(values); i++ {
		s[values[i]] = struct{}{}
	}
	return s
}

// Delete  从集合中删除1-n个元素，元素可以是任意类型
func (s Set[T]) Delete(values ...T) Set[T] {
	for _, value := range values {
		delete(s, value)
	}
	return s
}

// Clear 清空Set，返回空set
func (s Set[T]) Clear() Set[T] {
	for k := range s {
		delete(s, k)
	}
	return s
}

// Iterate 迭代器-传入一个方法，对集合中的每个元素执行该方法
func (s Set[T]) Iterate(it func(T)) {
	for v := range s {
		it(v)
	}
}

// Len 计算Set的容量
func (s Set[T]) Len() int {
	return len(s)
}

// Has 检查元素是否存在
func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

// String 实现string接口，优化打印输出
func (s Set[T]) String() string {
	return fmt.Sprint(s.ToSlice())
}

// ToSlice 将set转成slice，方便打印或输出
func (s Set[T]) ToSlice() []T {
	keys := make([]T, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}

func (s Set[T]) Clone() Set[T] {
	set := make(Set[T])
	for v := range s {
		set.Add(v)
	}
	return set
}

// Union 合并2个集合
func (s Set[T]) Union(other Set[T]) Set[T] {
	set := s.Clone()
	for v := range other {
		set.Add(v)
	}
	return set
}

// Intersection 交集
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	set := make(Set[T])
	s.Iterate(func(value T) {
		if other.Has(value) {
			set.Add(value)
		}
	})
	return set
}
