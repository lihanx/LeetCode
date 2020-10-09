package main

/*
559. N叉树的最大深度
https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree/
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// 1. 递归
func maxDepthRecursive(root *Node) int {
    if root == nil {
    	return 0
    } else if len(root.Children) == 0 {
    	return 1
    } else {
    	var depth int
    	for i := 0; i < len(root.Children); i++ {
    		d := maxDepth(root.Children[i])
    		if d > depth {
    			depth = d
    		}
    	}
    	return depth + 1
    }
}

// 2. 广度优先遍历
func maxDepthTraversal(root *Node) int {
	if root == nil {
		return 0
	}
	queue := []*Node{root}
	depth := 0
	for len(queue) > 0 {
		level := make([]*Node, 0)
		for i := 0; i < len(queue); i++ {
			node := queue[i]
			for _, c := range node.Children {
				level = append(level, c)
			}
		}
		depth++
		queue = level
	}
	return depth
}