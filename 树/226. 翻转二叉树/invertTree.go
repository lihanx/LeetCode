package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


/*
226. 翻转二叉树
https://leetcode-cn.com/problems/invert-binary-tree/submissions/
*/

func invertTree(root *TreeNode) *TreeNode {
	// 终止条件
	if root == nil {
		return root
	}
	// 镜像操作
	root.Left, root.Right = root.Right, root.Left
	// 逐级深入
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}