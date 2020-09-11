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
剑指 Offer 27. 二叉树的镜像
https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/

同 226. 翻转二叉树
https://leetcode-cn.com/problems/invert-binary-tree/submissions/
*/

func mirrorTree(root *TreeNode) *TreeNode {
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