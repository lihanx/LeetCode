package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


// 1. 递归
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return t1
	}
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val = t1.Val + t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}


// 2. 迭代
/*
迭代的方式有以下思路：

	1. 首先按照深度优先的方式对两树同时进行遍历，在遍历的过程中对节点进行合并
	2. 在遍历过程中，节点有以下几种组合情况：
		(1) node1 == nil && node2 == nil -> nil
		(2) node1 == nil && node2 != nil -> node1 = node2
		(3) node1 != nil && node2 == nil -> node1
		(4) node1 != nil && node2 != nil -> node1.Val += node2.Val

	3. 针对几种组合情况，实际上对应节点中任意一个为 nil，都不需要再继续向下遍历，只需要将子树放置到t1树对应位置即可

那么，同时对两树进行深度优先遍历，可以将两树对应节点放在同一个数组结构内（Slice），令他们成对入栈出栈。
这样就避免了额外维护对应关系，也更便于比较或计算。

同样的一点注意：不要尝试直接写出最简代码，根据自己最直接的思路完成代码后，再考虑简化逻辑
*/

func mergeTreesTraversal(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	// 首先初始化栈，将两树的根节点成对入栈
	stack := make([][]*TreeNode, 0)
	stack = append(stack, []*TreeNode{t1, t2})
	// 遍历过程中，自顶向下对节点进行操作
	for len(stack) > 0 {
		// 取出栈顶节点对
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 两节点均不为nil
		p[0].Val += p[1].Val
		// t1节点为nil时，将t2节点覆盖过来（即将子树整个放置过来）
		// 注意这里一定要 continue 跳过后续操作，否则会导致节点被计算两次
		// p[0].Left == nil && p[1].Left != nil
		// p[0].Left = p[1].Left
		// p[0].Left != nil && p[1].Left != nil
		// p[0].Left, p[1].Left -> stack
		if p[0].Left == nil && p[1].Left != nil {
			p[0].Left = p[1].Left
			continue
		}
		if p[0].Right == nil && p[1].Right != nil {
			p[0].Right = p[1].Right
			continue
		}
		// 这两种情况是不需要任何操作的
		if p[1].Left == nil && p[0].Left != nil {
			continue
		}
		if p[1].Right == nil && p[0].Right != nil {
			continue
		}
		
		// 这里保证放入栈的节点对均不为nil，所以在上方合并节点值时不需要再做额外的判断
		if p[0].Left != nil && p[1].Left != nil {
			stack = append(stack, []*TreeNode{p[0].Left, p[1].Left})
		}
		if p[0].Right != nil && p[1].Right != nil {
			stack = append(stack, []*TreeNode{p[0].Right, p[1].Right})
		}
	}
	return t1
}