package binarytree

import (
	"fmt"
	"go-dsa/common"
	"testing"
)

var at = NewAVLTree().(*AVLTree)

func init() {
	at.Add(common.Int(13))
	at.Add(common.Int(14))
	at.Add(common.Int(15))
	at.Add(common.Int(12))
	at.Add(common.Int(11))
	at.Add(common.Int(17))
	at.Add(common.Int(16))
	at.Add(common.Int(8))
	at.Add(common.Int(9))
	at.Add(common.Int(1))
}

func TestAVLTree_Add(t *testing.T) {
	LevelOrderTraversal(at.GetRoot(), func(node Node) bool {
		fmt.Println(node.GetEle())
		return true
	})

	fmt.Println(IsAVLTree(at))
}

func TestAVLTree_Remove(t *testing.T) {
	at.Remove(common.Int(1))
	at.Remove(common.Int(8))
	LevelOrderTraversal(at.GetRoot(), func(node Node) bool {
		fmt.Println(node.GetEle())
		return true
	})
	fmt.Println(IsAVLTree(at))
	at.Add(common.Int(1))
	at.Add(common.Int(8))
	LevelOrderTraversal(at.GetRoot(), func(node Node) bool {
		fmt.Println(node.GetEle())
		return true
	})
	fmt.Println(IsAVLTree(at))
}
