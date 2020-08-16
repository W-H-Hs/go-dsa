package 链表

// https://leetcode-cn.com/problems/middle-of-the-linked-list/

//func middleNode(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//	tmp := make([]*ListNode, 0)
//	cur := head
//	for cur != nil {
//		tmp = append(tmp, cur)
//		cur = cur.Next
//	}
//	if len(tmp)&1 == 1 {
//		return tmp[len(tmp)/2+1]
//	}
//	return tmp[len(tmp)/2]
//}
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	count := 0
	cur := head
	for cur != nil {
		count++
		cur = cur.Next
	}
	count /= 2
	for count >= 0 {
		head = head.Next
		count--
	}
	return head
}
