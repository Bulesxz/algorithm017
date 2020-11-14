#!/usr/bin/env python

# 剑指 Offer 68 - II. 二叉树的最近公共祖先
# https://leetcode-cn.com/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof/


from stack import Stack
from struct import TreeNode

class Solution(object):
    def lowestCommonAncestor(self, root, p, q):
        """
        :type root: TreeNode
        :type p: TreeNode
        :type q: TreeNode
        :rtype: TreeNode
        """
        if root == None:
            return None
        
        if p == root or q == root:
            return root
        left = self.lowestCommonAncestor(root.left,p,q)
        right = self.lowestCommonAncestor(root.right,p,q)

        if  left == None and right == None:
            return root
        if left!=None:
            return left
        if right!=None:
            return right
        return None

