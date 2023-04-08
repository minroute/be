package structure

import (
	"testing"
)

// cmd: go test -v int_slice_test.go int_slice.go -test.run TestIsIntSlice
func TestIsIntSlice(t *testing.T) {
	type v []int
	v1 := &v{1, 2, 3}
	if IsSlice(v1) {
		t.Error("kit.IsSlice error")
	}

	v2 := &IntSlice{1, 2, 3, 4}
	if !IsIntSlice(v2) {
		t.Error("IsIntSlice error")
	}

	if IsIntSlice(&[]int{1, 2, 3}) {
		t.Error("IsIntSlice error", &[]int{1, 2, 3})
	}

}

// 测试转换
// cmd: go test -v int_slice_test.go int_slice.go -test.run TestToIntSlice
func TestToIntSlice(t *testing.T) {
	v := ToIntSlice([]int{1, 2, 3})
	if !IsIntSlice(v) {
		t.Error("ToIntSlice error")
	}

	v2 := ToIntSliceFromPtrAddr(&[]int{1, 2, 3})
	if !IsIntSlice(v2) {
		t.Error("ToIntSliceFromPtr error")
	}

	temp := []int{1, 2, 3}
	v3 := IntSlice(temp)
	if !IsIntSlice(&v3) {
		t.Error("ToIntSliceFromPtr error")
	}

}

func TestIntSlice(t *testing.T) {

}
