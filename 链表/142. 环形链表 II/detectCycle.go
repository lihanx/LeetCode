package main

/*
142. 环形链表 II
https://leetcode-cn.com/problems/linked-list-cycle-ii/
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	fast, slow := head, head
	// 第一次相遇
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	// 开始计数
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}