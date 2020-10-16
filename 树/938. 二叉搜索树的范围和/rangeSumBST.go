package main

/*
938. 二叉搜索树的范围和
https://leetcode-cn.com/problems/range-sum-of-bst/
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
首先由于BST(Binary Search Tree)的特点，任一节点的Left < Right，要利用这一特性。
然后，遍历考虑深度优先、广度优先。
如果使用广度优先的方式，BST的特性将无法利用，必然需要遍历所有节点。
那么考虑使用深度优先，利用BST特点，减少不必要的节点访问。

利用迭代的方式，优先遍历左子树，在比较值大小的过程中，存在以下几种情况：
1. node.Val < L < R
2. L <= node.Val <= R
3. node.Val > R > L

当 1 时, 可能存在 L < node.Right.Val < R
当 2 时, 可能存在 L < node.Child.Val < R
当 3 时, 可能存在 L < node.Left.val < R


通过合并情况，得到以下条件:
1. 当 node.Val < R 时，需要继续遍历 node.Right
2. 当 node.Val > L 时, 需要继续遍历 node.Left

同时，对每一个满足范围的节点进行累加，可以得到最终结果。
*/

func rangeSumBST(root *TreeNode, L int, R int) int {
	var res int
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right != nil && node.Val <= R {
			stack = append(stack, node.Right)
		}
		if node.Left != nil && node.Val >= L {
			stack = append(stack, node.Left)
		}
		if node.Val >= L && node.Val <= R {
			res += node.Val
		}
	}
	return res
}