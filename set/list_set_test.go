package set

import (
	"fmt"
	"testing"
)

var listSet = NewListSet()

func TestListSet_Add(t *testing.T) {
	listSet.Add(10)
	listSet.Add(10)
	listSet.Add(11)
	listSet.Add(11)
	listSet.Add(12)
	listSet.Add(12)
	fmt.Println(listSet.Size())
	listSet.Traversal(Printer{})
}
