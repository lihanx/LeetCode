package main

/*
1290. 二进制链表转整数
https://leetcode-cn.com/problems/convert-binary-number-in-a-linked-list-to-integer/
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 利用 十进制 转 二进制的逆向运算
func getDecimalValue(head *ListNode) int {
	if head == nil {
		return 0
	}
	var ans int
	cur := head
	for cur != nil {
		ans = ans * 2 + cur.Val
		cur = cur.Next
	}
	return ans
}