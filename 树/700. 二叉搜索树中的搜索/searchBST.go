package main

/*
700. 二叉搜索树中的搜索
https://leetcode-cn.com/problems/search-in-a-binary-search-tree/
*/


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
利用BST(Binary Search Tree)节点值排列特点
对树结构进行深度优先遍历
*/

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		if node.Val == val {
			return node
		}
		if node.Val < val {
			stack = append(stack, node.Right)
		}
		if node.Val > val {
			stack = append(stack, node.Left)
		}
	}
	return nil
}