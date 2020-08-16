package binarytree

import (
	"fmt"
	"go-dsa/common"
	"testing"
)

func init() {
	bst.Add(common.Int(100))
	bst.Add(common.Int(50))
	bst.Add(common.Int(150))
	bst.Add(common.Int(25))
	bst.Add(common.Int(24))
	bst.Add(common.Int(75))
	bst.Add(common.Int(125))
	bst.Add(common.Int(151))
	bst.Add(common.Int(74))
	bst.Add(common.Int(80))
}

func TestPreorderTraversal(t *testing.T) {
	PreorderTraversal(bst.root, func(node Node) bool {
		fmt.Println(node.GetEle())
		return true
	})
}

func TestInorderTraversal(t *testing.T) {
	InorderTraversal(bst.root, func(node Node) bool {
		fmt.Println(node.GetEle())
		return true
	})
}

func TestPostorderTraversal(t *testing.T) {
	PostorderTraversal(bst.root, func(node Node) bool {
		fmt.Println(node.GetEle())
		return true
	})
}

func TestLevelOrderTraversal(t *testing.T) {
	LevelOrderTraversal(bst.root, func(node Node) bool {
		fmt.Println(node.GetEle())
		return true
	})
}

func TestGetPreviousNode(t *testing.T) {
	fmt.Println(GetPreviousNode(bst.root))
}

func TestGetNextNode(t *testing.T) {
	fmt.Println(GetNextNode(bst.root))
}

func TestGetHeight(t *testing.T) {
	fmt.Println(GetHeight(bst.root))
}

func TestIsCompleteBinaryTree(t *testing.T) {
	fmt.Println(IsCompleteBinaryTree(bst.root))
	bst.Remove(common.Int(74))
	bst.Remove(common.Int(80))
	fmt.Println(IsCompleteBinaryTree(bst.root))
}
