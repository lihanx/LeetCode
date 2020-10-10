package main

/*
237. 删除链表中的节点
https://leetcode-cn.com/problems/delete-node-in-a-linked-list/
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


// 由于只给定目标节点，对原链表和前面的节点一无所知，需要原地删除
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
	return
}