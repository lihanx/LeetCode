package main


/*
637. 二叉树的层平均值
https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 简单层序遍历，在遍历过程中增加求平均值
func averageOfLevels(root *TreeNode) []float64 {
	ans := make([]float64, 0)
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		var sum int
		level := make([]*TreeNode, 0)
		for i := 0; i < len(queue); i++ {
			sum += queue[i].Val

			if queue[i].Left != nil {
				level = append(level, queue[i].Left)
			}
			if queue[i].Right != nil {
				level = append(level, queue[i].Right)
			}
		}
		ans = append(ans, float64(sum)/float64(len(queue)))
		queue = level
	}
	return ans
}