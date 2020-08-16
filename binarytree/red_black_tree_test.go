package binarytree

import (
	"fmt"
	"go-dsa/common"
	"testing"
)

var rbt = NewRedBlackTree()
var nodes []common.Int

func init() {
	nodes = []common.Int{
		11, 57, 53, 83, 24, 68, 16, 18, 73, 33, 51, 44, 60, 10, 64, 70, 86, 25, 55, 20, 97, 63, 3, 79, 48, 27, 81, 22,
		8, 87, 76, 62, 91, 38, 12, 28, 17, 6, 50, 30, 69, 54, 45, 35, 14, 52, 49, 29, 42, 84, 96, 82, 90, 98, 19, 99,
		2, 43, 72, 58, 93, 66, 65, 94, 15, 78, 85, 59, 32, 47, 36, 100, 13, 1, 4, 34, 75, 80, 23, 95, 41, 89, 31, 92,
		46, 7, 21, 37, 9, 26, 74, 56, 67, 77, 61,
	}
	for _, v := range nodes {
		rbt.Add(v)
	}
}

func TestRedBlackTree_Add(t *testing.T) {
	LevelOrderTraversal(rbt.GetRoot(), func(node Node) bool {
		fmt.Println(node.GetEle(), node.(*RedBlackTreeNode).color)
		return true
	})
}

func TestRedBlackTree_Remove(t *testing.T) {
	for i := 0; i < len(nodes); i++ {
		rbt.Remove(nodes[i])
	}
	LevelOrderTraversal(rbt.GetRoot(), func(node Node) bool {
		fmt.Println(node.GetEle(), node.(*RedBlackTreeNode).color)
		return true
	})
}
