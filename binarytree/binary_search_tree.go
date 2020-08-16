package binarytree

import (
	"go-dsa/common"
)

type BinarySearchTreeNode struct {
	ele    common.ComparableElement
	left   *BinarySearchTreeNode
	right  *BinarySearchTreeNode
	parent *BinarySearchTreeNode
}

func (bn *BinarySearchTreeNode) GetParent() Node {
	return bn.parent
}

func (bn *BinarySearchTreeNode) SetParent(node Node) {
	if node == nil {
		bn.parent = nil
	} else {
		bn.parent = node.(*BinarySearchTreeNode)
	}
}

func (bn *BinarySearchTreeNode) GetLeft() Node {
	return bn.left
}

func (bn *BinarySearchTreeNode) SetLeft(node Node) {
	if node == nil {
		bn.left = nil
	} else {
		bn.left = node.(*BinarySearchTreeNode)
	}
}

func (bn *BinarySearchTreeNode) GetRight() Node {
	return bn.right
}

func (bn *BinarySearchTreeNode) SetRight(node Node) {
	if node == nil {
		bn.right = nil
	} else {
		bn.right = node.(*BinarySearchTreeNode)
	}
}

func (bn *BinarySearchTreeNode) GetEle() common.ComparableElement {
	return bn.ele
}

func (bn *BinarySearchTreeNode) SetEle(element common.ComparableElement) {
	bn.ele = element
}

type BinarySearchTree struct {
	*AbstractBinaryTree
}

func NewBinarySearchTree() BinaryTree {
	bst := &BinarySearchTree{}
	bst.AbstractBinaryTree = &AbstractBinaryTree{BinaryTree: bst}
	return bst
}

func (bst *BinarySearchTree) createNode(
	element common.ComparableElement,
	parent, left, right Node,
) Node {
	node := &BinarySearchTreeNode{ele: element}
	if parent != nil {
		node.parent = parent.(*BinarySearchTreeNode)
	}
	if left != nil {
		node.left = left.(*BinarySearchTreeNode)
	}
	if right != nil {
		node.right = right.(*BinarySearchTreeNode)
	}
	return node
}

func (bst *BinarySearchTree) Remove(element common.ComparableElement) (Node, error) {
	return bst.remove(element, nil)
}
