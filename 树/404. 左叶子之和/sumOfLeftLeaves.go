package main

/*
404. 左叶子之和
https://leetcode-cn.com/problems/sum-of-left-leaves/
*/


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
**/


func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

// 1. 递归 深度优先
func sumOfLeftLeavesDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root)
}

func dfs(node *TreeNode) int {
	var ans int
	if node.Left != nil {
		if isLeaf(node.Left) {
			ans += node.Left.Val
		} else {
			ans += dfs(node.Left)
		}
	}
	if node.Right != nil && !isLeaf(node.Right) {
		ans += dfs(node.Right)
	}
	return ans
}


// 2. 队列 广度优先
func sumOfLeftLeavesBFS(root *TreeNode) int {
	var ans int
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Left != nil {
			if isLeaf(node.Left) {
				ans += node.Left.Val
			} else {
				q = append(q, node.Left)
			}
		}
		if node.Right != nil && !isLeaf(node.Right) {
			q = append(q, node.Right)
		}
	}
	return ans
}