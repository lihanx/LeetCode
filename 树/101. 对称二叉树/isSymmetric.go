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
1. 迭代
迭代的思路，主要目标是，想办法在层序遍历的过程中进行比较
对于判断整个树是不是互为镜像（对称），可以看作是判断左右子树是否镜像
那么左右子树互为镜像的标准是：
    
    left.Left == right.Right && left.Right == right.Left

那么，在两两比较的过程中，根据树节点的不同，又可分为以下几种情况：
    
    1. node1 == nil && node2 == nil, 结论是相等
    2. (node1 == nil && node2 != nil) || (node1 != nil && node2 == nil) 结论是不相等
    3. node1 != nil && node2 != nil
        3.1 node1.Val == node2.Val 结论是相等
        3.2 node1.Val != node2.Val 结论是不等

在代码实现中，可以根据逻辑对几种情况的判断进行简并
    
    1. 同为 nil 的情况，无法获取子节点，直接判断为相等，继续下一轮比较，首先过滤
    2. 单个为 nil 的情况，直接判为不等，返回 false，然后过滤
    3. 均不为 nil 时，首先判断是否不等。若不等，直接返回 false；
       若相等，需要继续获取相应的子节点，放入队列中等待下一轮比较

这里值得注意的一点是，不应该在最开始就尝试写出最优的代码，这样会平添难度；
最先考虑的问题应该是，如何梳理自己的思路，并按照思路准确地实现代码；
当代码可以得到正确的结果后，再考虑优化逻辑，简化代码；
*/
func isSymmetricTraversal(root *TreeNode) bool {
    if root == nil {
        return true
    }
    queue := []*TreeNode{root.Left, root.Right}
    for len(queue) > 0 {
        u, v := queue[0], queue[1]
        queue = queue[2:]
        if u == nil && v == nil {
            continue
        }
        if u == nil || v == nil {
            return false
        }
        if u.Val != v.Val {
            return false
        }
        queue = append(queue, u.Left)
        queue = append(queue, v.Right)
        
        queue = append(queue, u.Right)
        queue = append(queue, v.Left)
    }
    return true
}

/*
2. 递归
递归的方法，从思路上可能会更容易理解。
当我们发现，左右子树互为镜像的条件是，所有的节点都满足：

    left.Left == right.Right && left.Right == right.Left

对与非 root 节点来说，每个节点都有两个子节点，那么比较就需要进行两次
对于 root 节点来说，只有两个子节点，那么只需要进行一次
在递归的过程中，就可以从 root.Left 和 root.Right 开始比较, 减少一次重复判断
*/
func isSymmetricRecursive(root *TreeNode) bool {
    if root == nil {
        return true
    }    
    if isEqual(root.Left, root.Right) {
        return true
    }
    return false
}

func isEqual(node1, node2 *TreeNode) bool {
    if node1 == nil && node2 == nil {
        return true
    }
    if node1 == nil || node2 == nil {
        return false
    }
    if node1.Val != node2.Val {
        return false
    }
    return isEqual(node1.Left, node2.Right) && isEqual(node1.Right, node2.Left)
}