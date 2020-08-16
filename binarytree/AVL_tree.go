package binarytree

import (
	"go-dsa/common"
	"math"
	"reflect"
)

type AVLTreeNode struct {
	ele    common.ComparableElement
	left   *AVLTreeNode
	right  *AVLTreeNode
	parent *AVLTreeNode
	height int
}

func (atn *AVLTreeNode) GetParent() Node {
	return atn.parent
}

func (atn *AVLTreeNode) SetParent(node Node) {
	if node == nil {
		atn.parent = nil
	} else {
		atn.parent = node.(*AVLTreeNode)
	}
}

func (atn *AVLTreeNode) GetLeft() Node {
	return atn.left
}

func (atn *AVLTreeNode) SetLeft(node Node) {
	if node == nil {
		atn.left = nil
	} else {
		atn.left = node.(*AVLTreeNode)
	}
}

func (atn *AVLTreeNode) GetRight() Node {
	return atn.right
}

func (atn *AVLTreeNode) SetRight(node Node) {
	if node == nil {
		atn.right = nil
	} else {
		atn.right = node.(*AVLTreeNode)
	}
}

func (atn *AVLTreeNode) GetEle() common.ComparableElement {
	return atn.ele
}

func (atn *AVLTreeNode) SetEle(element common.ComparableElement) {
	atn.ele = element
}

func (atn *AVLTreeNode) calculateHeight() (int, int) {
	leftHeight := 0
	rightHeight := 0
	if atn.left != nil {
		leftHeight = atn.left.height
	}
	if atn.right != nil {
		rightHeight = atn.right.height
	}
	return leftHeight, rightHeight
}

func (atn *AVLTreeNode) calculateBalanceFactor() int {
	leftHeight, rightHeight := atn.calculateHeight()
	return leftHeight - rightHeight
}

func (atn *AVLTreeNode) updateHeight() {
	leftHeight, rightHeight := atn.calculateHeight()
	// 当前节点的高度等于左子树和右子树中最大值加一。
	atn.height = int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}

// 返回左子节点和右子节点中高度更高的子节点，如果两者高度相同，
// 则返回该节点与父节点同方向的子节点。
func (atn *AVLTreeNode) tallerChild() *AVLTreeNode {
	leftHeight, rightHeight := atn.calculateHeight()
	if leftHeight > rightHeight {
		return atn.left
	} else if leftHeight < rightHeight {
		return atn.right
	} else {
		// 默认返回与父节点同方向的子节点。
		if atn.parent.left == atn {
			return atn.left
		} else {
			return atn.right
		}
	}
}

type AVLTree struct {
	*AbstractBinaryTree
}

func NewAVLTree() BinaryTree {
	at := &AVLTree{}
	at.AbstractBinaryTree = &AbstractBinaryTree{BinaryTree: at}
	return at
}

func (at *AVLTree) createNode(
	element common.ComparableElement,
	parent, left, right Node,
) Node {
	// 新插入AVL树的节点一定是叶子节点，其高度一定为1。
	node := &AVLTreeNode{ele: element, height: 1}
	if parent != nil {
		node.parent = parent.(*AVLTreeNode)
	}
	if left != nil {
		node.left = left.(*AVLTreeNode)
	}
	if right != nil {
		node.right = right.(*AVLTreeNode)
	}

	return node
}

func (at *AVLTree) isBalanced(node *AVLTreeNode) bool {
	factor := node.calculateBalanceFactor()
	return -1 <= factor && factor <= 1
}

// 参数node一定是从新添加的节点往上遍历，遇到的第一个不平衡节点，
// 即高度最低的不平衡节点。
func (at *AVLTree) recoverBalance(node *AVLTreeNode) {
	// 添加节点导致失衡的时候，必定代表被添加的子树的高度大于另外
	// 一棵子树。
	child := node.tallerChild()
	grandChild := child.tallerChild()
	if node.left == child {
		if child.left == grandChild {
			// LL，进行右旋。
			at.rotate(node, false)
		} else {
			// LR，首先对node高度更高的子节点进行左旋，变成LL情况，
			// 然后对node进行右旋。
			at.rotate(child, true)
			at.rotate(node, false)
		}
	} else {
		if child.left == grandChild {
			// RL，首先对node高度更高的子节点进行右旋，变成RR情况，
			// 然后对node进行左旋。
			at.rotate(child, false)
			at.rotate(node, true)
		} else {
			// RR，进行左旋。
			at.rotate(node, true)
		}
	}
}

func (at *AVLTree) rotate(parent *AVLTreeNode, isLeft bool) {
	commonStep := func(child *AVLTreeNode) {
		if parent.parent != nil {
			if parent.parent.left == parent {
				parent.parent.left = child
				child.parent = parent.parent
			} else {
				parent.parent.right = child
				child.parent = parent.parent
			}
			parent.parent = child
		} else {
			// 失衡节点的父节点为空，代表失衡节点为根节点parent，旋转
			// 之后根节点不再是parent而是child，因此AVLTree的根节点
			// 属性也需要修改。
			child.parent = nil
			parent.parent = child
			at.root = child
		}

		// 更新旋转节点的高度。
		parent.updateHeight()
		child.updateHeight()
	}

	if isLeft {
		// 左旋。
		child := parent.right

		parent.right = child.left
		if parent.right != nil {
			// 右子树不是单个节点。
			parent.right.parent = parent
		}
		child.left = parent
		commonStep(child)
	} else {
		// 右旋。
		child := parent.left

		parent.left = child.right
		if parent.left != nil {
			// 左子树不是单个节点。
			parent.left.parent = parent
		}
		child.right = parent
		commonStep(child)
	}
}

func (at *AVLTree) recover(node Node) {
	cur := node
	for !reflect.ValueOf(cur).IsNil() {
		if at.isBalanced(cur.(*AVLTreeNode)) {
			// 这里恢复平衡的节点一定是不平衡节点的子节点。
			cur.(*AVLTreeNode).updateHeight()
		} else {
			// 找到第一个不平衡节点，让其恢复平衡后，整棵树都平衡了，
			// 此时不用继续往上遍历来恢复高度和寻找不平衡节点。
			at.recoverBalance(cur.(*AVLTreeNode))
			break
		}
		cur = cur.GetParent()
	}
}

func (at *AVLTree) Add(element common.ComparableElement) (Node, error) {
	node, err := at.AbstractBinaryTree.Add(element)
	at.recover(node)
	return node, err
}

func (at *AVLTree) Remove(element common.ComparableElement) (Node, error) {
	node, err := at.AbstractBinaryTree.remove(element, func(node Node) {
		at.recover(node)
	})
	return node, err
}

// 检查所有节点的平衡因子是否属于[-1, 1]。
func IsAVLTree(tree *AVLTree) bool {
	isAVLTree := true
	LevelOrderTraversal(tree.GetRoot(), func(node Node) bool {
		if !tree.isBalanced(node.(*AVLTreeNode)) {
			isAVLTree = false
			return false
		}
		return true
	})
	return isAVLTree
}
