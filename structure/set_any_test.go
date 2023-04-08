package structure

import (
	"encoding/json"
	"fmt"
	"testing"
)

type testStruct struct {
	Name string
	Age  int
}

func TestAnySet(t *testing.T) {

	// 类型
	fmt.Println("添加内置类型：数字/字符串/nil/bool")
	s := AnySet(true).Add("1").Add(true).Add(nil).Add("4", "5")
	fmt.Printf("set:%v, len:%v\n", s, s.Len())

	// 注意类型安全默认是关闭的·
	fmt.Printf("添加结构体,结构体类型只能是boolean, numeric, string, pointer, channel, interface,否则会出错\n")
	s2 := AnySet().Add(&testStruct{Name: "test", Age: 1})
	fmt.Printf("set2:%v, len:%v\n", s2, s2.Len())
	s2.Iterate(func(v interface{}) {
		fmt.Printf("set2 iterate:%+v\n", v)
	})

	fmt.Printf("添加限定长度的数组，不能添加未知长度的Slice，如需Slice，请使用SliceSet\n")
	s2.Add([3]string{"1", "2", "3"})
	s2.Add([3]int{1, 2, 3})
	fmt.Printf("set:%v, len:%v\n", s2, s2.Len())

	// MarshalJSON
	r, e := json.Marshal(s2) // or 	// r, e := s2.MarshalJSON()
	if e != nil {
		t.Error(e)
	}
	fmt.Printf("JSON:\n")
	fmt.Printf("set2 to json string:%v\n", string(r))
	fmt.Printf("set2 invoke JsonString:%v\n", s2.JsonString())

	// Output：
	// 添加内置类型：数字/字符串/nil/bool
	// set:[1 true <nil> 4 5], len:5
	// 添加结构体,结构体类型只能是boolean, numeric, string, pointer, channel, interface,否则会出错
	// set2:[0x1400012a078], len:1
	// set2 iterate:&{Name:test Age:1}
	// 添加限定长度的数组，不能添加未知长度的Slice，如需Slice，请使用SliceSet
	// set:[0x1400012a078 [1 2 3] [1 2 3]], len:3
	// JSON:
	// set2 to json string:[{"Name":"test","Age":1},["1","2","3"],[1,2,3]]
	// set2 invoke JsonString:[{"Name":"test","Age":1},["1","2","3"],[1,2,3]]
}
