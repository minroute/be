package structure

import "testing"

func TestStrAnyMapPtr(t *testing.T) {
	s := &StrAnyMap{"a": "b", "c": &[]string{"1", "2"}}
	if !IsStrAnyMap(s) {
		t.Error("IsStrAnyMap error")
	}
}

func TestStrAnyMap(t *testing.T) {
	s := StrAnyMap{"a": "b", "c": &[]string{"1", "2"}}
	if !IsStrAnyMap(s) {
		t.Error("IsStrAnyMap error")
	}
}
