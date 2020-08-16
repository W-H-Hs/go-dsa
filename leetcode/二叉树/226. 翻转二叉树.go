package 二叉树

/**
 * Definition for a binary tree_tmp node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type queue struct {
	array []interface{}
}

func (q *queue) enqueue(ele interface{}) {
	q.array = append(q.array, ele)
}
func (q *queue) dequeue() interface{} {
	e := q.array[len(q.array)-1]
	newArr := make([]interface{}, len(q.array)-1)
	copy(newArr, q.array[:len(q.array)-1])
	q.array = newArr
	return e
}
func (q *queue) len() int {
	return len(q.array)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func invertTree(root *TreeNode) *TreeNode {
//	if root == nil {
//		return nil
//	}
//	//q := list.NewCycleQueue(5)
//	q := queue{}
//	q.enqueue(root)
//	var tmp *TreeNode
//	for q.len() != 0 {
//		node := q.dequeue().(*TreeNode)
//		tmp = node.Left
//		node.Left = node.Right
//		node.Right = tmp
//		if node.Left != nil {
//			q.enqueue(node.Left)
//		}
//		if node.Right != nil {
//			q.enqueue(node.Right)
//		}
//	}
//	return root
//}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	invertTree(root.Left)
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	invertTree(root.Left)

	return root
}
