package 链表

// 给定一个链表，判断链表中是否有环。
// https://leetcode-cn.com/problems/linked-list-cycle

func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
