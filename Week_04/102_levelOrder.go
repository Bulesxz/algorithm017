package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//就是BFS  队列
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		level++
		res = append(res, []int{})
		curSize := len(queue)
		tmp := queue
		queue = []*TreeNode{}
		for i := 0; i < curSize; i++ {
			top := tmp[i]
			res[level-1] = append(res[level-1], top.Val)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
	}
	return res
}

//dfs 需要记录层级
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	dfsLevelOrder(root, 1, &res)
	return res
}
func dfsLevelOrder(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}
	if len((*res)) < level {
		*res = append(*res, []int{})
	}
	(*res)[level-1] = append((*res)[level-1], root.Val)

	if root.Left != nil {
		dfsLevelOrder(root.Left, level+1, res)
	}
	if root.Right != nil {
		dfsLevelOrder(root.Right, level+1, res)
	}
}
