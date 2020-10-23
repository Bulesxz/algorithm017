package main

import (
	"fmt"
	"math"
)

/*
124. 二叉树中的最大路径和
https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
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
1. 此函数的返回值 为该节点的最大贡献值，而不是最大路径和
*/
func dfsMaxPathSum(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	//1. 左子树最大路径和
	maxLeft := dfsMaxPathSum(root.Left, max)
	fmt.Println("maxLeft", maxLeft, max)
	if maxLeft < 0 {
		//不选此子树
		maxLeft = 0
	}
	//2. 右子树最大路径和
	maxRight := dfsMaxPathSum(root.Right, max)
	fmt.Println("maxRight", maxRight, max)
	if maxRight < 0 {
		//不选此子树
		maxRight = 0
	}
	//3. 节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
	curMax := maxLeft + maxRight + root.Val
	if curMax > *max {
		*max = curMax
	}
	//4. 返回节点的最大贡献值
	if maxLeft > maxRight {
		return maxLeft + root.Val
	} else {
		return maxRight + root.Val
	}

}
func maxPathSum(root *TreeNode) int {
	max := math.MinInt32
	dfsMaxPathSum(root, &max)
	return max
}
