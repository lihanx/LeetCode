package main

import (
	"math"
)

/*
剑指 Offer 55 - II. 平衡二叉树
https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/

与主站 110 题相同
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
**/



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
	if math.Abs(float64(lh-rh)) <= 1 {
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