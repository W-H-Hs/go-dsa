package binarytree

import (
	"errors"
	"go-dsa/common"
	"reflect"
)

type Node interface {
	GetParent() Node
	SetParent(Node)
	GetLeft() Node
	SetLeft(Node)
	GetRight() Node
	SetRight(Node)
	GetEle() common.ComparableElement
	SetEle(common.ComparableElement)
}

type BinaryTree interface {
	GetRoot() Node
	Size() int
	IsEmpty() bool
	Clear()
	Add(common.ComparableElement) (Node, error)
	Remove(common.ComparableElement) (Node, error)
	IsContain(common.ComparableElement) bool

	createNode(common.ComparableElement, Node, Node, Node) Node
}

type AbstractBinaryTree struct {
	BinaryTree
	size int
	root Node
}

func (a *AbstractBinaryTree) GetRoot() Node {
	return a.root
}

func (a *AbstractBinaryTree) Size() int {
	return a.size
}

func (a *AbstractBinaryTree) IsEmpty() bool {
	return a.size == 0
}

func (a *AbstractBinaryTree) Clear() {
	a.size = 0
	a.root = nil
}

func (a *AbstractBinaryTree) GetSpecificNode(
	element common.ComparableElement,
	exec func(node Node),
) Node {
	cur := a.root
	for {
		switch element.CompareTo(cur.GetEle()) {
		case -1:
			if reflect.ValueOf(cur.GetLeft()).IsNil() && exec != nil {
				exec(cur)
				break
			}
			cur = cur.GetLeft()
			if reflect.ValueOf(cur).IsNil() {
				return nil
			}
		case 0:
			return cur
		case 1:
			if reflect.ValueOf(cur.GetRight()).IsNil() && exec != nil {
				exec(cur)
				break
			}
			cur = cur.GetRight()
			if reflect.ValueOf(cur).IsNil() {
				return nil
			}
		}
	}
}

func (a *AbstractBinaryTree) Add(element common.ComparableElement) (Node, error) {
	if element == nil {
		return nil, errors.New("element is nil")
	}

	newNode := a.createNode(element, nil, nil, nil)

	if a.root == nil {
		a.root = newNode
		a.size++
		return newNode, nil
	}

	node := a.GetSpecificNode(element, func(node Node) {
		if element.CompareTo(node.GetEle()) == -1 {
			node.SetLeft(newNode)
		} else {
			node.SetRight(newNode)
		}
		newNode.SetParent(node)
	})
	if node != nil {
		node.SetEle(element)
	}
	a.size++
	return newNode, nil
}

// Node的度只可能为0或1。
func (a *AbstractBinaryTree) removeNode(node Node) {
	if reflect.ValueOf(node.GetLeft()).IsNil() {
		// 左子节点为空。
		if reflect.ValueOf(node.GetRight()).IsNil() {
			// 右子节点为空。
			// 删除度为0的节点。
			if reflect.ValueOf(node.GetParent()).IsNil() {
				a.root = nil
			} else {
				if node.GetParent().GetLeft() == node {
					node.GetParent().SetLeft(nil)
				} else {
					node.GetParent().SetRight(nil)
				}
			}
		} else {
			// 右子节点不为空。
			// 删除拥有右子节点的节点。
			if reflect.ValueOf(node.GetParent()).IsNil() {
				node.GetRight().SetParent(nil)
				a.root = node.GetRight()
			} else {
				if node.GetParent().GetLeft() == node {
					node.GetParent().SetLeft(node.GetRight())
				} else {
					node.GetParent().SetRight(node.GetRight())
				}
				node.GetRight().SetParent(node.GetParent())
			}
		}
	} else {
		// 左子节点不为空。
		if reflect.ValueOf(node.GetRight()).IsNil() {
			// 右子节点为空。
			// 删除拥有左子节点的节点。
			if reflect.ValueOf(node.GetParent()).IsNil() {
				node.GetLeft().SetParent(nil)
				a.root = node.GetLeft()
			} else {
				if node.GetParent().GetLeft() == node {
					node.GetParent().SetLeft(node.GetLeft())
				} else {
					node.GetParent().SetRight(node.GetLeft())
				}
				node.GetLeft().SetParent(node.GetParent())
			}

		}
	}
}

func (a *AbstractBinaryTree) remove(
	element common.ComparableElement,
	execute func(node Node),
) (Node, error) {
	node := a.GetSpecificNode(element, nil)
	if node == nil {
		return nil, errors.New("element not in tree")
	}

	// 删除度为2的节点，可以使用中序遍历的前驱节点或后继节点来取代该节点的值，
	// 然后删除替代节点，这里使用后继节点。替代节点的度必为0或1。
	if !reflect.ValueOf(node.GetLeft()).IsNil() &&
		!reflect.ValueOf(node.GetRight()).IsNil() {
		nextNode := GetNextNode(node)
		node.SetEle(nextNode.GetEle())
		a.removeNode(nextNode)
		if execute != nil {
			execute(nextNode)
		}
	} else {
		a.removeNode(node)
		if execute != nil {
			execute(node)
		}
	}

	return node, nil
}

func (a *AbstractBinaryTree) IsContain(element common.ComparableElement) bool {
	node := a.GetSpecificNode(element, nil)
	return node == nil
}
