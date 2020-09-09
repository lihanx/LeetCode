package main

/*
590. N叉树的后序遍历
https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */


// 1. 递归解法
func postorder(root *Node) []int {
	res := make([]int, 0)
    if root == nil {
    	return res
    }
    lrv(root, &res)
    return res
}

func lrv(node *Node, res *[]int) {
	for _, c := range node.Children {
		lrv(c, res)
	}
	*res = append(*res, node.Val)
}


// 2. 迭代解法
func postorderTraversal(root *Node) []int {
	
}