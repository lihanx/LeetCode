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


/* 
2. 迭代解法 
迭代解法这里其实没有直接的方式可以完成
通过对比前序遍历和后序遍历的结果，可以发现，后序遍历的结果可以通过对前序遍历进行改造获得
*/
func postorderTraversal(root *Node) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		res = append(res, node.Val)
		stack = stack[:len(stack)-1]
		for _, c := range node.Children {
			stack = append(stack, c)
		}
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}