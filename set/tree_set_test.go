package set

import (
	"go-dsa/common"
	"testing"
)

var treeSet = NewTreeSet()

func TestTreeSet_Add(t *testing.T) {
	treeSet.Add(common.Int(10))
	treeSet.Add(common.Int(10))
	treeSet.Add(common.Int(11))
	treeSet.Add(common.Int(11))
	treeSet.Add(common.Int(12))
	treeSet.Add(common.Int(12))
	treeSet.Traversal(Printer{})
}
