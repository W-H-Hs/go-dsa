package set

import (
	"go-dsa/binarytree"
	"go-dsa/common"
)

type TreeSet struct {
	tree *binarytree.RedBlackTree
}

func NewTreeSet() *TreeSet {
	return &TreeSet{
		binarytree.NewRedBlackTree().(*binarytree.RedBlackTree),
	}
}

func (t *TreeSet) Size() int {
	return t.tree.Size()
}

func (t *TreeSet) IsEmpty() bool {
	return t.tree.IsEmpty()
}

func (t *TreeSet) Clear() {
	t.tree.Clear()
}

func (t *TreeSet) IsContain(ele interface{}) bool {
	switch ele.(type) {
	case common.ComparableElement:
		return t.tree.IsContain(ele.(common.ComparableElement))
	default:
		panic("element must can be comparable")
	}
}

func (t *TreeSet) Add(ele interface{}) (interface{}, error) {
	switch ele.(type) {
	case common.ComparableElement:
		return t.tree.Add(ele.(common.ComparableElement))
	default:
		panic("element must can be comparable")
	}
}

func (t *TreeSet) Remove(ele interface{}) (interface{}, error) {
	switch ele.(type) {
	case common.ComparableElement:
		return t.tree.Remove(ele.(common.ComparableElement))
	default:
		panic("element must can be comparable")
	}
}

func (t *TreeSet) Traversal(visitor Visitor) {
	if visitor == nil {
		return
	}

	binarytree.InorderTraversal(t.tree.GetRoot(), func(node binarytree.Node) bool {
		visitor.Visit(node.GetEle())
		return true
	})
}
