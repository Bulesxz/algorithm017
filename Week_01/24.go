package main

/*
24. 两两交换链表中的节点
https://leetcode-cn.com/problems/swap-nodes-in-pairs/
*/

//递归版本
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first := head
	second := first.Next
	first.Next = swapPairs(second.Next)
	second.Next = first
	return second
}

//循环遍历版本

func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ret := head.Next
	pre := &ListNode{Next: head}
	for head != nil && head.Next != nil {
		first := head
		second := first.Next
		third := second.Next

		//second->first->third
		//pre->second
		second.Next = first
		first.Next = third
		pre.Next = second

		//pre->first
		//head->third
		pre = first
		head = third
	}

	return ret
}
