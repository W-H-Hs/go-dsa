package list

import (
	"fmt"
	"go-dsa/common"
	"testing"
)

var sl = NewSkipList()

func init() {
	_, _ = sl.Put(common.Int(1), 1)
	_, _ = sl.Put(common.Int(2), 2)
	_, _ = sl.Put(common.Int(3), 3)
}

func TestSkipList_Remove(t *testing.T) {
	fmt.Println(sl.Get(common.Int(1)))
	fmt.Println(sl.Get(common.Int(2)))
	fmt.Println(sl.Get(common.Int(3)))
	removedNode, err := sl.Remove(common.Int(1))
	if err != nil {
		panic(err)
	}
	fmt.Println(removedNode)
	fmt.Println(sl.Get(common.Int(1)))
	fmt.Println(sl.Get(common.Int(2)))
	fmt.Println(sl.Get(common.Int(3)))
}
