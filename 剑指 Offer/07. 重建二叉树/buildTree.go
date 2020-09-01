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
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]


1.根据preorder可知，preorder[0] 为 root节点
2.根据root节点在inorder结果中的位置，可将inorder结果分为leftSubTree, root, rightSubTree
3.inorder的SubTree元素数量，可以用于将preorder结果区分为左右子树
4.重复1
 
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
    posMap := make(map[int]int)
    for idx, val := range inorder {
        posMap[val] = idx
    }
    return recurBuild(0, 0, len(inorder)-1, &posMap, preorder, inorder)
}

func recurBuild(preRoot, inLeft, inRight int, pMap *map[int]int, preorder, inorder []int) *TreeNode {
    if in_left > in_right {
        return (*TreeNode)(nil)
    }
    root := new(TreeNode)
    root.Val = preorder[preRoot]
    pos := (*pMap)[preorder[preRoot]]
    root.Left = recurBuild(0, inLeft, pos, preorder, inorder)
}