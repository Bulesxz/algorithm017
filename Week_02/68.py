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
        times=0
        st = Stack()
        st.push(root)
        while not st.is_empty():
            item = st.top()
            if item!=None and (item.val == p.val or item.val == q.val):
                times  = times + 1
                break
            if item.right !=None:
                st.push(item.right)
            if item.left !=None:
                st.push(item.left)


