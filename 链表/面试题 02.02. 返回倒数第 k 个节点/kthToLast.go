package main

/*
面试题 02.02. 返回倒数第 k 个节点
https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci/
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 双指针
func kthToLast(head *ListNode, k int) int {
	target, cur := head, head
	for i := 0; i < k; i++ {
		cur = cur.Next
	}
	for cur != nil {
		target = target.Next
		cur = cur.Next
	}
	return target.Val
}