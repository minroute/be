package structure

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	set := make(Set[string])
	set.Add("1")
	set.Add("2")
	set.Add("3")
	set.Add("4", "5")
	fmt.Printf("set:%v, len:%v\n", set, len(set))
	set.Add("4")
	fmt.Printf("add 4, set:%v\n", set)
	set.Delete("2")
	fmt.Printf("del 2, set:%v\n", set)
	fmt.Printf("contain 4, result is :%v\n", set.Has("4"))
	fmt.Printf("contain 44, result is :%v\n", set.Has("44"))

	set2 := set.Clone()
	fmt.Printf("set Clone:%v,origin set is :%v\n", set2, set)

	set2.Add("6", "7")
	set2 = set2.Union(set)
	fmt.Printf("set2 union set:%v\n", set2)
	fmt.Printf("set2 intersection set:%v\n", set2.Intersection(set))

	set.Clear()
	fmt.Printf("set clear:%v\n", set)

	// Output:
	// set:[1 2 3 4 5], len:5
	// add 4, set:[1 2 3 4 5]
	// del 2, set:[1 3 4 5]
	// contain 4, result is :true
	// contain 44, result is :false
	// set Clone:[1 3 4 5],origin set is :[1 3 4 5]
	// set2 union set:[1 3 4 5 6 7]
	// set2 intersection set:[1 3 4 5]
	// set clear:[]
}
