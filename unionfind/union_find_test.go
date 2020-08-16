package unionfind

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	quickFind  = NewUnionFind(10, false)
	quickUnion = NewUnionFind(10, true)
)

func test(unionFind UnionFind) {
	unionFind.Union(0, 1)
	unionFind.Union(1, 2)
	unionFind.Union(2, 3)
	unionFind.Union(4, 5)
	unionFind.Union(5, 6)
	unionFind.Union(6, 7)
	fmt.Println(unionFind.IsSame(0, 1))
	fmt.Println(unionFind.IsSame(0, 3))
	fmt.Println(unionFind.IsSame(4, 7))
	fmt.Println(unionFind.IsSame(0, 7))
	unionFind.Union(0, 7)
	fmt.Println(unionFind.IsSame(0, 7))
}

func TestQuickFind(t *testing.T) {
	test(quickFind)
}

func TestUnionFind(t *testing.T) {
	test(quickUnion)
}

func tt(s interface{}) {
	ss := reflect.ValueOf(s)
	sss := ss.Slice(0, ss.Len())
	fmt.Println(sss.Index(0))
	fmt.Println(sss.Index(1))
	fmt.Println(sss.Index(2))
	fmt.Println(reflect.TypeOf(sss.Index(0).Interface()))
}

func demo(s interface{}) {
	sv := reflect.ValueOf(s)
	svs := sv.Slice(0, sv.Len())
	for i := 0; i < svs.Len(); i++ {
		e := svs.Index(i).Interface()
		switch e.(type) {
		case string:
			fmt.Println("string", e)
		case int:
			fmt.Println("int", e)
		}
	}
}

func Test(t *testing.T) {
	s1 := []int{1, 2, 3}
	demo(s1)
	s2 := []string{"a", "b", "c"}
	demo(s2)
}
