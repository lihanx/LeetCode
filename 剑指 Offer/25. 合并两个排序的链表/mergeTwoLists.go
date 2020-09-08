package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
剑指 Offer 25. 合并两个排序的链表
https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/
*/ 

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{-1, nil}
    last := dummy
    ptr1, ptr2 := l1, l2
    for ptr1 != nil && ptr2 != nil {
        if ptr1.Val <= ptr2.Val {
            last.Next = ptr1
            ptr1 = ptr1.Next
        } else {
            last.Next = ptr2
            ptr2 = ptr2.Next
        }
        last = last.Next
    }
    if ptr1 != nil {
        last.Next = ptr1
    }
    if ptr2 != nil {
        last.Next = ptr2
    }
    return dummy.Next
}