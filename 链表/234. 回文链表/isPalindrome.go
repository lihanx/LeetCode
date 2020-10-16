package main

/*
234. 回文链表
https://leetcode-cn.com/problems/palindrome-linked-list/
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	// 定位中点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 翻转后半段链表
	var prev *ListNode
	cur := slow
	for cur != nil {
		slow = slow.Next
		cur.Next = prev
		prev = cur
		cur = slow
	}
	// 比较
	// prev, head
	cur = head
	for cur != nil && prev != nil {
		if prev.Val != cur.Val {
			return false
		}
		prev = prev.Next
		cur = cur.Next
	}
	return true
}