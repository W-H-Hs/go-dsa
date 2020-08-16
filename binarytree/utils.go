package binarytree

import (
	"go-dsa/list"
	"reflect"
)

// 前序遍历
func PreorderTraversal(root Node, execute func(Node) bool) {
	if root == nil {
		return
	}

	s := list.NewStack()
	s.Push(root)
	for s.Size() != 0 {
		node := s.Pop().(Node)
		if !execute(node) {
			return
		}
		if !reflect.ValueOf(node.GetRight()).IsNil() {
			s.Push(node.GetRight())
		}
		if !reflect.ValueOf(node.GetLeft()).IsNil() {
			s.Push(node.GetLeft())
		}
	}
}

func pushLeftChild(node Node, s *list.Stack) {
	cur := node
	for !reflect.ValueOf(cur).IsNil() {
		s.Push(cur)
		cur = cur.GetLeft()
	}
}

// 中序遍历
func InorderTraversal(root Node, execute func(Node) bool) {
	if root == nil {
		return
	}

	s := list.NewStack()
	pushLeftChild(root, s)
	for s.Size() != 0 {
		node := s.Pop().(Node)
		if !execute(node) {
			return
		}
		if !reflect.ValueOf(node.GetRight()).IsNil() {
			pushLeftChild(node.GetRight(), s)
		}
	}
}

// 后序遍历
func PostorderTraversal(root Node, execute func(Node) bool) {
	if root == nil {
		return
	}

	s := list.NewStack()
	length := list.NewStack()
	pushLeftChild(root, s)
	for s.Size() != 0 {
		node := s.Top().(Node)
		if !reflect.ValueOf(node.GetRight()).IsNil() {
			// 若s.Size() == length.Top().(int)，则代表node节点的右子树
			// 已经被遍历过了。
			if s.Size() != length.Top().(int) {
				length.Push(s.Size())
				pushLeftChild(node.GetRight(), s)
			} else {
				length.Pop()
				s.Pop()
				if !execute(node) {
					return
				}
			}
		} else {
			s.Pop()
			if !execute(node) {
				return
			}
		}
	}
}

// 层序遍历
func LevelOrderTraversal(root Node, execute func(node Node) bool) {
	if root == nil {
		return
	}

	q := list.NewCycleQueue(5)
	q.EnQueue(root)
	for q.Len() != 0 {
		node := q.DeQueue().(Node)
		if !execute(node) {
			return
		}
		if !reflect.ValueOf(node.GetLeft()).IsNil() {
			q.EnQueue(node.GetLeft())
		}
		if !reflect.ValueOf(node.GetRight()).IsNil() {
			q.EnQueue(node.GetRight())
		}
	}
}

func isRightChild(parent, child Node) bool {
	return reflect.DeepEqual(parent.GetRight(), child)
}

// 中序遍历的前一个节点。
func GetPreviousNode(node Node) Node {
	if node == nil {
		return nil
	}

	left := node.GetLeft()
	if !reflect.ValueOf(left).IsNil() {
		// 当左子树不为空的时候，前驱节点为左子树中最底层
		// 的最后一个节点。
		cur := left
		for !reflect.ValueOf(cur.GetRight()).IsNil() {
			cur = cur.GetRight()
		}
		return cur
	} else {
		if reflect.ValueOf(node.GetParent()).IsNil() {
			// 左子树为空的根节点，没有前驱节点。
			return nil
		}

		// 左子树为空，则从当前节点往上回溯的所有祖先节点中，某节点
		// 与其子节点第一次满足子节点是这个节点的右子节点时，该节点
		// 为当前节点的前驱节点。
		cur := node
		parent := cur.GetParent()
		for !reflect.ValueOf(parent).IsNil() {
			if isRightChild(parent, cur) {
				return parent
			}
			cur = parent
			parent = parent.GetParent()
		}
	}

	// 上述条件都不满足，为最底层第一个节点。
	return nil
}

func isLeftChild(parent, child Node) bool {
	return reflect.DeepEqual(parent.GetLeft(), child)
}

// 中序遍历的后一个节点。
func GetNextNode(node Node) Node {
	if node == nil {
		return nil
	}

	right := node.GetRight()
	if !reflect.ValueOf(right).IsNil() {
		// 当右子树不为空的时候，后继节点为右子树中最底层
		// 的第一个节点。
		cur := right.GetLeft()
		if reflect.ValueOf(cur).IsNil() {
			// 若right的左子节点为空，则right即为后继节点。
			return right
		}
		for !reflect.ValueOf(cur.GetLeft()).IsNil() {
			// 当cur的左子节点为空的时候，cur即为右子树最
			// 底层的第一个节点。
			cur = cur.GetLeft()
		}
		return cur
	} else {
		if reflect.ValueOf(right.GetParent()).IsNil() {
			// 右子树为空的根节点，没有后继节点。
			return nil
		}

		// 右子树为空，则从当前节点往上回溯的所有祖先节点中，
		// 若某节点与其子节点第一次满足子节点是某节点的左子
		// 节点时，该节点为后继节点。
		cur := node
		parent := cur.GetParent()
		for !reflect.ValueOf(parent).IsNil() {
			if isLeftChild(parent, cur) {
				return parent
			}
			cur = parent
			parent = parent.GetParent()
		}
	}

	// 上诉所有情况都不满足，则为最后一个节点。
	return nil
}

func GetHeight(root Node) int {
	if root == nil {
		return 0
	}

	height := 0
	cur := 1
	next := 0
	LevelOrderTraversal(root, func(node Node) bool {
		if !reflect.ValueOf(node.GetLeft()).IsNil() {
			next++
		}
		if !reflect.ValueOf(node.GetRight()).IsNil() {
			next++
		}
		cur--
		if cur == 0 {
			cur = next
			next = 0
			height++
		}
		return true
	})

	return height
}

// 判断是否是完全二叉树。
func IsCompleteBinaryTree(root Node) bool {
	if root == nil {
		return false
	}

	var (
		isComplete = true
		isLeaf     = false
	)

	LevelOrderTraversal(root, func(node Node) bool {
		if isLeaf {
			// 当遍历到的节点为叶节点时，左右子节点必须为空，否则不为
			// 完全二叉树。
			if !reflect.ValueOf(node.GetLeft()).IsNil() ||
				!reflect.ValueOf(node.GetRight()).IsNil() {
				isComplete = false
				return false
			}
		}

		if reflect.ValueOf(node.GetLeft()).IsNil() {
			if !reflect.ValueOf(node.GetRight()).IsNil() {
				// 左子节点为空，右子节点不为空，必定不为完全二叉树。
				isComplete = false
				return false
			}

			// 左右子节点为空，接下来遍历的节点均为叶节点。
			isLeaf = true
		} else {
			if reflect.ValueOf(node.GetRight()).IsNil() {
				// 左子节点不为空右子节点为空，接下来遍历的节点均为叶节点。
				isLeaf = true
			}
		}

		return true
	})

	return isComplete
}
