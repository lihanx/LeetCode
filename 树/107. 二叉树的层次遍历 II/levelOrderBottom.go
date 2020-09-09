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
107. 二叉树的层次遍历 II
https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/submissions/

树是没有办法自底向上层序遍历的，拿到 root 节点，必然以 root 为起点
那么自然没有一步到位的方法可以直接获取到自底向上层序遍历的结果
所以只能考虑先自顶向下进行层序遍历，然后将结果翻转为逆序，则得到自底向上的层序遍历结果
*/

func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	// 初始化队列
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		// 创建层级队列和结果slice
		level := make([]*TreeNode, 0)
		levelRes := make([]int, 0)
		// 遍历队列生成新的队列，同时添加层级结果
		for i := 0; i < len(queue); i++ {
			levelRes = append(levelRes, queue[i].Val)
			if queue[i].Left != nil {
				level = append(level, queue[i].Left)	
			}
			if queue[i].Right != nil {
				level = append(level, queue[i].Right)
			}
		}
		res = append(res, levelRes)
		queue = level
	}
	// 根据正序结果翻转，获得逆序结果
	i, j := 0, len(res) - 1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}