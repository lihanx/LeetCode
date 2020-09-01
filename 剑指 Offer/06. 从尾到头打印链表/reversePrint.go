package main

func reversePrint(head *ListNode) []int {
    stack := make([]int, 0)
    res := make([]int, 0)
    curr := head
    for curr != nil {
        stack = append(stack, curr.Val)
        curr = curr.Next
    }
    for i := len(stack)-1; i >=0; i-- {
        res = append(res, stack[i])
    }
    return res
}