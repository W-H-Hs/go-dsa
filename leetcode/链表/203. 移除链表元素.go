package 链表

// https://leetcode-cn.com/problems/remove-linked-list-elements/

func removeElements(head *ListNode, val int) *ListNode {
	cur := head
	for cur != nil {
		if cur.Val == val {
			head = head.Next
			cur = head
		} else {
			if cur.Next == nil {
				break
			}
			if cur.Next.Val == val {
				cur.Next = cur.Next.Next
			} else {
				cur = cur.Next
			}
		}
	}
	return head
}
