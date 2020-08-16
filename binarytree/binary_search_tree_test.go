package binarytree

import (
	"fmt"
	"go-dsa/common"
	"testing"
)

var bst = NewBinarySearchTree().(*BinarySearchTree)

func init() {
	nodes := []common.Int{100, 50, 150, 25, 24, 75, 125, 151, 74, 80}
	for _, v := range nodes {
		bst.Add(v)
	}
}

func TestBinarySearchTree_Add(t *testing.T) {
	LevelOrderTraversal(bst.root, func(node Node) bool {
		fmt.Println(node)
		return true
	})
}

func TestBinarySearchTree_Remove(t *testing.T) {
	bst.Remove(common.Int(24))
	//bst.remove(com(100))
	LevelOrderTraversal(bst.root, func(node Node) bool {
		fmt.Println(node)
		//if node.GetEle() == com(24){
		//	fmt.Println(node)
		//	fmt.Println(node.GetParent())
		//}
		return true
	})
}
