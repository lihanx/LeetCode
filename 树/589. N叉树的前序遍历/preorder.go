package main

/*
589. N叉树的前序遍历
https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/submissions/
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */


/*
1. 递归解法
*/
func preorderRecursive(root *Node) []int {
	res := make([]int, 0)
    if root == nil {
    	return res
    }
    vlr(root, &res)
    return res
}


func vlr(node *Node, res *[]int) {
	*res = append(*res, node.Val)
	for _, c := range node.Children {
		vlr(c, res)
	}
}


/*
2. 迭代解法
*/
func preorderTraversal(root *Node) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for i := len(node.Children)-1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
		res = append(res, node.Val)
	}
	return res
}