package main

/*
94. 二叉树的中序遍历
https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
*/


// 1. 递归
func inorderTraversal1(root *TreeNode) []int {
    res := make([]int, 0)
    if root == nil {
        return res
    }
    lvr(root, &res)
    return res
}

func lvr(node *TreeNode, res *[]int) {
    if node.Left != nil {
        lvr(node.Left, res)
    }
    *res = append(*res, node.Val)
    if node.Right != nil {
        lvr(node.Right, res)
    }
}


// 2. 迭代
func inorderTranversal2(root *TreeNode) []int {
    res := make([]int, 0)
    stack := make([]*TreeNode, 0)
    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        // 这里不能使用 := 声明赋值
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res = append(res, root.Val)
        root = root.Right
    }
    return res
}