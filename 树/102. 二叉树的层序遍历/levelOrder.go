package main

import (
    "fmt"
)

/*
102.二叉树的层序遍历
https://leetcode-cn.com/problems/binary-tree-level-order-traversal
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func levelOrder(root *TreeNode) [][]int {
    res := make([][]int, 0)
    // 临界值
    if root == nil {
        return res
    }
    // 初始化队列
    queue := []*TreeNode{root}
    // 定式循环
    for i := 0; len(queue) > 0; i++ {
        // 创建本层数组
        res = append(res, []int{})
        // 创建本层子节点队列
        tmpQueue := make([]*TreeNode, 0)
        // 遍历主队列
        for j := 0; j < len(queue); j++ {
            node := queue[j]
            res[i] = append(res[i], node.Val)
            if node.Left != nil {
                tmpQueue = append(tmpQueue, node.Left)
            }
            if node.Right != nil {
                tmpQueue = append(tmpQueue, node.Right)
            }
        }
        queue = tmpQueue
    }
    return res
}