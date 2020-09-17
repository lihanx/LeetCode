package main

import (
	"math"
)

/*
110. 平衡二叉树
https://leetcode-cn.com/problems/balanced-binary-tree/
*/


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
**/


/*
递归 - 自底向上计算
*/

func isBalanced(root *TreeNode) bool {
	if height(root) != -1 {
		return true
	}
	return false
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	lh := height(node.Left)
	if lh == -1 {
		return -1
	}
	rh := height(node.Right)
	if rh == -1 {
		return -1
	}
	if math.Abs(float64(lh - rh)) <= 1 {
		return max(lh, rh) + 1
	} else {
		return -1
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}