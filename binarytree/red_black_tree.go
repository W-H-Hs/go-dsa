package binarytree

import (
	"go-dsa/common"
	"reflect"
)

// 1.节点是red或black；
// 2.根节点是black；
// 3.叶子节点（外部节点，空节点）都是black；
// 4.red节点的子节点都是black、red节点的父节点都是black、
//  从根节点到叶节点的所有路径都不能包含连续两个red节点；
// 5.任一节点到叶节点的所有路径都包含相同数目的black节点；

type NodeColor bool

func (nc NodeColor) String() string {
	if nc {
		return "RED"
	} else {
		return "BLACK"
	}
}

const (
	Red   NodeColor = true
	Black NodeColor = false
)

type RedBlackTreeNode struct {
	ele    common.ComparableElement
	left   *RedBlackTreeNode
	right  *RedBlackTreeNode
	parent *RedBlackTreeNode
	color  NodeColor
}

func (rn *RedBlackTreeNode) GetParent() Node {
	return rn.parent
}

func (rn *RedBlackTreeNode) SetParent(node Node) {
	if node == nil {
		rn.parent = nil
	} else {
		rn.parent = node.(*RedBlackTreeNode)
	}
}

func (rn *RedBlackTreeNode) GetLeft() Node {
	return rn.left
}

func (rn *RedBlackTreeNode) SetLeft(node Node) {
	if node == nil {
		rn.left = nil
	} else {
		rn.left = node.(*RedBlackTreeNode)
	}
}

func (rn *RedBlackTreeNode) GetRight() Node {
	return rn.right
}

func (rn *RedBlackTreeNode) SetRight(node Node) {
	if node == nil {
		rn.right = nil
	} else {
		rn.right = node.(*RedBlackTreeNode)
	}
}

func (rn *RedBlackTreeNode) GetEle() common.ComparableElement {
	return rn.ele
}

func (rn *RedBlackTreeNode) SetEle(element common.ComparableElement) {
	rn.ele = element
}

func (rn *RedBlackTreeNode) getSibling() Node {
	if rn.parent == nil {
		return nil
	}

	if rn == rn.parent.left {
		return rn.parent.right
	} else {
		return rn.parent.left
	}
}

func (rn *RedBlackTreeNode) isLeft() bool {
	if rn.parent.left == rn {
		return true
	} else {
		return false
	}
}

type RedBlackTree struct {
	*AbstractBinaryTree
}

func NewRedBlackTree() BinaryTree {
	rbt := &RedBlackTree{}
	rbt.AbstractBinaryTree = &AbstractBinaryTree{BinaryTree: rbt}
	return rbt
}

func (rbt *RedBlackTree) createNode(
	element common.ComparableElement,
	parent, left, right Node,
) Node {
	node := &RedBlackTreeNode{ele: element}
	if rbt.root == nil {
		node.color = Black
	} else {
		node.color = Red
	}
	if parent != nil {
		node.parent = parent.(*RedBlackTreeNode)
	}
	if left != nil {
		node.left = left.(*RedBlackTreeNode)
	}
	if right != nil {
		node.right = right.(*RedBlackTreeNode)
	}
	return node
}

func (rbt *RedBlackTree) dye(node Node, color NodeColor) Node {
	if node == nil {
		return nil
	}

	node.(*RedBlackTreeNode).color = color
	return node
}

func (rbt *RedBlackTree) dyeToRed(node Node) {
	rbt.dye(node, Red)
}

func (rbt *RedBlackTree) dyeToBlack(node Node) {
	rbt.dye(node, Black)
}

func (rbt *RedBlackTree) isBlack(node Node) bool {
	if reflect.ValueOf(node).IsNil() {
		return true
	} else {
		if node.(*RedBlackTreeNode).color == Black {
			return true
		} else {
			return false
		}
	}
}

func (rbt *RedBlackTree) Add(element common.ComparableElement) (Node, error) {
	node, err := rbt.AbstractBinaryTree.Add(element)
	if !reflect.ValueOf(node).IsNil() {
		parent := node.GetParent()
		if reflect.ValueOf(parent).IsNil() {
			return node, err
		}
		if parent.(*RedBlackTreeNode) != rbt.root && parent.(*RedBlackTreeNode).color == Red {
			// 插入节点的父节点不为根节点，且父节点为红色。
			rbt.recoverBalanceForAdd(node)
		}
	}
	return node, err
}

// 新添加的节点最好默认为red，这样能满足1、2、3、5，但性质4不一定满足。
// 如果添加的是根节点，直接染为black即可。
// 节点添加可以分为parent为黑、parent为红两种情况，继续往下又可以分为12种情况：
// 一、parent为黑，此时不需要修复红黑树的性质：
// 		1.parent只有右子节点，添加到parent的左子节点处；
// 		2.parent只有左子节点，添加到parent的右子节点处；
// 		3.parent没有子节点，添加到parent的左子节点处；
// 		4.parent没有子节点，添加到parent的右子节点处；
// 二、parent为红，此时需要修复红黑树的性质：
// 		5.parent存在sibling且parent为左子节点，添加到parent的左子节点处；
// 		6.parent存在sibling且parent为左子节点，添加到parent的右子节点处；
// 		7.parent存在sibling且parent为右子节点，添加到parent的左子节点处；
// 		8.parent存在sibling且parent为右子节点，添加到parent的右子节点处；
// 		9.parent不存在sibling且parent为左子节点，添加到parent的左子节点处；
// 		10.parent不存在sibling且parent为左子节点，添加到parent的右子节点处；
// 		11.parent不存在sibling且parent为右子节点，添加到parent的左子节点处；
// 		12.parent不存在sibling且parent为右子节点，添加到parent的右子节点处；
func (rbt *RedBlackTree) recoverBalanceForAdd(node Node) {
	// 在旋转中，谁要变成父节点，就将谁染为黑色，谁要变成子节点，就将谁染为红色。
	if reflect.ValueOf(node.GetParent().(*RedBlackTreeNode).getSibling()).IsNil() ||
		node.GetParent().(*RedBlackTreeNode).getSibling().(*RedBlackTreeNode).color == Black {
		// 无叔父节点或叔父节点为黑色（两者等价），且父节点为红色。
		if node.(*RedBlackTreeNode).isLeft() {
			if node.GetParent().(*RedBlackTreeNode).isLeft() {
				// 插入节点是父节点的左子节点，父节点是祖父节点的左子节点，即LL。
				// 进行的操作是将父节点染为黑色，祖父节点染为红色，然后对祖父节
				// 点进行右旋转。
				rbt.rotate(node.GetParent().GetParent(), false, false, func(parent, child Node) {
					rbt.dyeToRed(parent)
					rbt.dyeToBlack(child)
				})
			} else {
				// 插入节点是父节点的左子节点，父节点是祖父节点的右子节点，即LR。
				// 进行的操作是首先将插入节点染为黑色因为插入节点要变为父节点，
				// 对父节点进行右旋，然后将祖父节点染为红色因为祖父节点要变为子
				// 节点，对祖父节点进行左旋。
				parent := node.GetParent()
				grandpa := parent.GetParent()
				// 注意旋转关系，是先旋转parent然后旋转grandpa。
				rbt.rotate(parent, false, false, func(parent, child Node) {
					rbt.dyeToRed(parent)
					rbt.dyeToBlack(child)
				})
				rbt.rotate(grandpa, true, false, func(parent, child Node) {
					rbt.dyeToRed(parent)
					rbt.dyeToBlack(child)
				})
			}
		} else {
			if node.GetParent().(*RedBlackTreeNode).isLeft() {
				// 插入节点是父节点的右子节点，父节点是祖父节点的左子节点，即RL。
				// 进行的操作是首先将插入节点染为黑色因为插入节点要变为父节点，
				// 对父节点进行左旋，然后将祖父节点染为红色因为祖父节点要变为子
				// 节点，对祖父节点进行右旋。
				parent := node.GetParent()
				grandpa := parent.GetParent()
				// 注意旋转关系，是先旋转parent然后旋转grandpa。
				rbt.rotate(parent, true, false, func(parent, child Node) {
					rbt.dyeToRed(parent)
					rbt.dyeToBlack(child)
				})
				rbt.rotate(grandpa, false, false, func(parent, child Node) {
					rbt.dyeToRed(parent)
					rbt.dyeToBlack(child)
				})
			} else {
				// 插入节点是父节点的右子节点，父节点是祖父节点的右子节点，即RR。
				// 进行的操作是将父节点染为黑色，祖父节点染为黑色，然后对祖父节
				// 点进行左旋转。
				rbt.rotate(node.GetParent().GetParent(), true, false, func(parent, child Node) {
					rbt.dyeToRed(parent)
					rbt.dyeToBlack(child)
				})
			}
		}
	} else {
		// 拥有叔父节点，且父节点与叔父节点一定是红色。
		// 首先将插入节点的父节点和叔父节点染为黑色，如果祖父节点为根节点，那么整棵树
		// 已经平衡，否则将祖父节点染为红色，如果祖父节点的父节点为黑色，那么整棵树已
		// 经平衡，否则将祖父节点作为新插入节点来进行恢复平衡。
		rbt.dyeToBlack(node.GetParent())
		rbt.dyeToBlack(node.GetParent().(*RedBlackTreeNode).getSibling())
		if !reflect.ValueOf(node.GetParent().GetParent().GetParent()).IsNil() {
			grandpa := node.GetParent().GetParent()
			rbt.dyeToRed(grandpa)
			if grandpa.GetParent().(*RedBlackTreeNode).color == Red {
				rbt.recoverBalanceForAdd(node.GetParent().GetParent())
			}
		}
	}
}

func (rbt *RedBlackTree) Remove(element common.ComparableElement) (Node, error) {
	node, err := rbt.AbstractBinaryTree.remove(element, func(node Node) {
		if node.(*RedBlackTreeNode).color == Red {
			// 删除节点为红色不需要平衡。
			return
		}

		if reflect.ValueOf(node.GetParent()).IsNil() {
			// 根节点不需要平衡。
			return
		}

		// node的度只是0或1。
		if reflect.ValueOf(node.GetLeft()).IsNil() {
			if reflect.ValueOf(node.GetRight()).IsNil() {
				// node无子节点，node为实际被删除节点。
				rbt.recoverBalanceForRemove(node, true)
			} else {
				// node有右子节点，右子节点为替代节点，即实际被删除节点。
				if node.GetRight().(*RedBlackTreeNode).color == Red {
					rbt.dyeToBlack(node.GetRight())
				} else {
					rbt.recoverBalanceForRemove(node.GetRight(), true)
				}
			}
		} else {
			// node有左子节点，左子节点为替代节点，即实际被删除节点。
			if node.GetLeft().(*RedBlackTreeNode).color == Red {
				rbt.dyeToBlack(node.GetLeft())
			} else {
				rbt.recoverBalanceForRemove(node.GetLeft(), true)
			}
		}
	})
	return node, err
}

// 节点删除可以分为节点为红、黑两种情况，如果是删除红色节点，那么不需要恢复平衡，
// 删除黑色节点分为以下情况：
// 1.兄弟节点为黑色且拥有一个红色子节点（不论父节点红黑）；
// 2.兄弟节点为黑色且拥有两个红色子节点（不论父节点红黑）；
// 3.兄弟节点为黑色且无红色子节点（区分父节点红黑）；
// 4.兄弟节点为红色；
// 兄弟节点为黑色的时候，注意先后顺序，即如果兄弟节点没有子节点才让父节点向下合并。
// 由于红黑树的性质，被删除节点为黑色时，如果兄弟节点为黑色，其要么拥有红色子节点
// 要么没有子节点。
// 参数isDeleted表示当前节点是否已经删除。
func (rbt *RedBlackTree) recoverBalanceForRemove(node Node, isDeleted bool) {
	removedNode := node.(*RedBlackTreeNode)
	var (
		isLeft bool
		// 不能通过getSibling()函数来获得兄弟节点，因为node为已经删除节点，父节点
		// 子节点处已经为空，如果调用getSibling()函数将会发生空指针异常。
		sibling *RedBlackTreeNode
	)
	if isDeleted {
		if reflect.ValueOf(removedNode.GetParent().GetLeft()).IsNil() {
			// 父节点的左子节点为空，代表removeNode为左子节点，因此sibling为右子节点。
			isLeft = true
			sibling = removedNode.GetParent().GetRight().(*RedBlackTreeNode)
		} else {
			// 父节点的右子节点为空，代表removeNode为右子节点，因此sibling为左子节点。
			isLeft = false
			sibling = removedNode.GetParent().GetLeft().(*RedBlackTreeNode)
		}
	} else {
		sibling = removedNode.getSibling().(*RedBlackTreeNode)
		if !sibling.isLeft() {
			isLeft = true
		}
	}

	if rbt.isBlack(sibling) {
		if rbt.isBlack(sibling.left) && rbt.isBlack(sibling.right) {
			if removedNode.parent.color == Black {
				// 3.2 被删除节点的父节点为黑色，将兄弟节点染红，当父节点不为根节点时
				// 将父节点作为被删除节点来进行递归；
				rbt.dyeToRed(sibling)
				if !reflect.ValueOf(node.GetParent().GetParent()).IsNil() {
					// 由于父节点实际并未被删除，因此isDeleted应该取false。
					rbt.recoverBalanceForRemove(node.GetParent(), false)
				}
			} else {
				// 3.1 被删除节点的父节点为红色，直接将兄弟节点染红，父节点染黑；
				rbt.dyeToBlack(node.GetParent())
				rbt.dyeToRed(sibling)
				return
			}
		}
		if !rbt.isBlack(sibling.left) && rbt.isBlack(sibling.right) {
			if !isLeft {
				// 1.1 被删除节点为右子节点且兄弟节点拥有红色左子节点，删除节点后父节点、兄弟
				// 节点、兄弟节点的左子节点形成LL情况，对父节点进行右旋，旋转完成后将新父节点
				// 的左右子节点都染黑，但需要注意的是这里父节点下沉被染红，但兄弟节点上浮继承
				// 父节点的原颜色而不是直接染黑；
				rbt.rotate(node.GetParent(), false, true, func(parent, child Node) {
					rbt.dye(child, parent.(*RedBlackTreeNode).color)
					rbt.dyeToBlack(parent)
				})
			} else {
				// 1.3 被删除节点为左子节点且兄弟节点拥有红色左子节点，删除节点后父节点、兄弟
				// 节点、兄弟节点的右子节点形成RL情况，对兄弟节点进行右旋对父节点进行左旋，旋
				// 转完成后将新父节点的左右子节点都染黑，但需要注意的是这里父节点下沉被染红，
				// 兄弟节点的左子节点上浮继承父节点的原颜色而不是直接染黑；
				rbt.rotate(sibling, false, false, func(parent, child Node) {})
				rbt.rotate(node.GetParent(), true, true, func(parent, child Node) {
					rbt.dye(child, parent.(*RedBlackTreeNode).color)
					rbt.dyeToBlack(parent)
				})
			}
		} else if rbt.isBlack(sibling.left) && !rbt.isBlack(sibling.right) {
			if !isLeft {
				// 1.2 被删除节点为右子节点且兄弟节点拥有红色右子节点，删除节点后父节点、兄弟
				// 节点、兄弟节点的右子节点形成LR情况，对兄弟节点进行左旋对父节点进行右旋，旋
				// 转完成后将新父节点的左右子节点都染黑，但需要注意的是这里父节点下沉被染红，
				// 兄弟节点的右子节点上浮继承父节点的原颜色而不是直接染黑；
				rbt.rotate(sibling, true, false, func(parent, child Node) {})
				rbt.rotate(node.GetParent(), false, true, func(parent, child Node) {
					rbt.dye(child, parent.(*RedBlackTreeNode).color)
					rbt.dyeToBlack(parent)
				})
			} else {
				// 1.4 被删除节点为左子节点且兄弟节点拥有红色右子节点，删除节点后父节点、兄弟
				// 节点、兄弟节点的左子节点形成RR情况，对父节点进行左旋，旋转后将新父节点的左
				// 右子节点都染黑，但需要注意的是这里父节点下沉被染红，但兄弟节点上浮继承父节
				// 点的原颜色而不是直接染黑；
				rbt.rotate(node.GetParent(), true, true, func(parent, child Node) {
					rbt.dye(child, parent.(*RedBlackTreeNode).color)
					rbt.dyeToBlack(parent)
				})
			}
			//} else if sibling.left != nil && sibling.right != nil || sibling.left.color == Black && sibling.right.color == Black {
		} else if !rbt.isBlack(sibling.left) && !rbt.isBlack(sibling.right) {
			if !isLeft {
				// 2.1 被删除节点为右子节点且兄弟节点拥有两个红色子节点，此时可以用LL或LR的
				// 情况处理，最好是LL，旋转染成后将新父节点的左右子节点都染黑；
				rbt.rotate(node.GetParent(), false, true, func(parent, child Node) {
					rbt.dye(child, parent.(*RedBlackTreeNode).color)
					rbt.dyeToBlack(parent)
				})
			} else {
				// 2.2 被删除节点为左子节点且兄弟节点拥有两个红色子节点，此时可以用RR或RL的
				// 情况处理，最好是RR，旋转染成后将新父节点的左右子节点都染黑；
				rbt.rotate(node.GetParent(), true, true, func(parent, child Node) {
					rbt.dye(child, parent.(*RedBlackTreeNode).color)
					rbt.dyeToBlack(parent)
				})
			}
		}
	} else {
		if isLeft {
			// 4.1 被删除节点为左子节点且兄弟节点为红色，父节点染为红色兄弟节点染为黑色，将父
			// 节点进行一次左旋，就将兄弟节点的子节点变为了兄弟节点（根据红黑树特性4和5，红色
			// 兄弟节点一定拥有黑色子节点），然后进行递归，此时情况变成了兄弟节点为黑色的情况。
			rbt.rotate(node.GetParent(), true, false, func(parent, child Node) {
				rbt.dyeToRed(parent)
				rbt.dyeToBlack(child)
			})
			if reflect.ValueOf(node.GetParent().GetLeft()).IsNil() {
				rbt.recoverBalanceForRemove(node, true)
			} else {
				rbt.recoverBalanceForRemove(node, false)
			}
		} else {
			// 4.2 被删除节点为右子节点且兄弟节点为红色，父节点染为红色兄弟节点染为黑色，将父
			// 节点进行一次右旋，就将兄弟节点的子节点变为了兄弟节点（根据红黑树特性4和5，红色
			// 兄弟节点一定拥有黑色子节点），然后进行递归，此时情况变成了兄弟节点为黑色的情况。
			rbt.rotate(node.GetParent(), false, false, func(parent, child Node) {
				rbt.dyeToRed(parent)
				rbt.dyeToBlack(child)
			})
			if reflect.ValueOf(node.GetParent().GetRight()).IsNil() {
				rbt.recoverBalanceForRemove(node, true)
			} else {
				rbt.recoverBalanceForRemove(node, false)
			}
		}
	}
}

func (rbt *RedBlackTree) rotate(node Node, isLeft, isNeedDying bool, dye func(parent, child Node)) {
	commonStep := func(child Node) {
		if !reflect.ValueOf(node.GetParent()).IsNil() {
			// node不是根节点。
			if node.(*RedBlackTreeNode).isLeft() {
				node.GetParent().SetLeft(child)
			} else {
				node.GetParent().SetRight(child)
			}
		} else {
			// node是根节点，旋转之后根节点已经发生了改变。
			rbt.root = child
		}
		child.SetParent(node.GetParent())
		node.SetParent(child)
	}
	var child Node
	if isLeft {
		// 左旋。
		child = node.GetRight()
		dye(node, child)
		if !reflect.ValueOf(child.GetLeft()).IsNil() {
			node.SetRight(child.GetLeft())
			child.GetLeft().SetParent(node)
		} else {
			node.SetRight(nil)
		}
		commonStep(child)
		child.SetLeft(node)
	} else {
		// 右旋。
		child = node.GetLeft()
		dye(node, child)
		if !reflect.ValueOf(child.GetRight()).IsNil() {
			node.SetLeft(child.GetRight())
			child.GetRight().SetParent(node)
		} else {
			node.SetLeft(nil)
		}
		commonStep(child)
		child.SetRight(node)
	}

	if isNeedDying {
		rbt.dyeToBlack(child.GetLeft())
		rbt.dyeToBlack(child.GetRight())
	}
}
