package main

/*
965. 单值二叉树
https://leetcode-cn.com/problems/univalued-binary-tree/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 广度优先遍历
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			if node.Left.Val != root.Val {
				return false
			} else {
				queue = append(queue, node.Left)
			}
		}
		if node.Right != nil {
			if node.Right.Val != root.Val {
				return false
			} else {
				queue = append(queue, node.Right)
			}
		}
	}
	return true
}