# -*- coding:utf-8

# 102.二叉树的层序遍历
# https://leetcode-cn.com/problems/binary-tree-level-order-traversal

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        res = []
        # 临界值判断
        if root is None:
            return res
        # 初始化队列
        q = [root]
        i = 0
        while len(q) > 0:
            tmp_q = []
            res.append([])
            for node in q:
                res[i].append(node.val)
                if node.left is not None:
                    tmp_q.append(node.left)
                if node.right is not None:
                    tmp_q.append(node.right)
            i += 1
            q = tmp_q
        return res