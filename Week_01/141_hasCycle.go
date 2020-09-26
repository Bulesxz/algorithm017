package main

/*
141. 环形链表
https://leetcode-cn.com/problems/linked-list-cycle/
*/

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next.Next
	for fast != nil {
		if fast == slow {
			return true
		} else {
			slow = slow.Next
			if fast.Next == nil {
				return false
			} else {
				fast = fast.Next.Next
			}
		}
	}
	return false
}
