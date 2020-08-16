package 链表

// https://leetcode-cn.com/problems/reverse-linked-list/
// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL

//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	newHead := reverseList(head.Next)
//	head.Next.Next = head
//	head.Next = nil
//	return newHead
//}
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead *ListNode
	for {
		tmp := head.Next
		head.Next = newHead
		newHead = head
		head = tmp
		if head == nil {
			break
		}
	}
	return newHead
}
