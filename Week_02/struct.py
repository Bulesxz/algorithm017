#!/user/bin/env python
#  -*- coding:utf-8 -*-

from stack import Stack

# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


#迭代版本
def pre_traversal(root,visit):
    if root == None:
        return
    if not isinstance(root,TreeNode):
        raise Exception("非法类型")
    st = Stack()
    st.push(root)
    while not st.is_empty():
        item = st.pop()
        visit(item)
        if item.right !=None:
            st.push(item.right)
        if item.left !=None:
            st.push(item.left)

#递归版本
def pre_traversal_recursive(root,visit):
    if root == None:
        return
    if not isinstance(root,TreeNode):
        raise Exception("非法类型")
    if root != None:
        visit(root)
    pre_traversal_recursive(root.left, visit)
    pre_traversal_recursive(root.right, visit)

if __name__ == "__main__":
    n1 = TreeNode(1)
    n2 = TreeNode(2)
    n3 = TreeNode(3)
    n4 = TreeNode(4)
    n5 = TreeNode(5)
    n6 = TreeNode(6)
    n7 = TreeNode(7)
    n7.left  = n6
    n7.right = n5
    n6.left  = n4
    n6.right = n3
    n5.left = n2
    n4.right = n1
    def f(x):
        print x.val
    pre_traversal(n7,f)

    pre_traversal_recursive(n7, f)


 
