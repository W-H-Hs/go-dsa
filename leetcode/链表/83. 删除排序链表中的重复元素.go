package 链表

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/

//func deleteDuplicates(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	slow := head
//	fast := head
//	for fast.Next != nil {
//		fast = fast.Next
//		if fast.Val == slow.Val && fast.Next == nil {
//			slow.Next = nil
//			break
//		}
//		if fast.Val != slow.Val {
//			slow.Next = fast
//			slow = fast
//		}
//	}
//	return head
//}

//func deleteDuplicates(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	current := head
//	for {
//		if current.Val == current.Next.Val {
//			current.Next = current.Next.Next
//			if current.Next == nil {
//				break
//			}
//		}
//		if current.Val != current.Next.Val {
//			current = current.Next
//			if current.Next == nil {
//				break
//			}
//		}
//	}
//	return head
//}
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
