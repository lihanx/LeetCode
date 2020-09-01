package main

/*
111. 二叉树的最小深度
https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/submissions/
*/


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    // 初始化队列
    queue := []*TreeNode{root}
    depth := 1
    // bfs
    for len(queue) > 0 {
        layer := make([]*TreeNode, 0)
        for j := 0; j < len(queue); j++ {
            node := queue[j]
            if node.Left == nil && node.Right == nil {
                return depth
            }
            if node.Left != nil {
                layer = append(layer, node.Left)
            }
            if node.Right != nil {
                layer = append(layer, node.Right)
            }
        }
        depth++
        queue = layer
    }
    return depth
}