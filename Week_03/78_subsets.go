package main

import (
	"fmt"
	"sort"
)

/*
78. 子集
https://leetcode-cn.com/problems/subsets/
*/

//回溯
func backtraceSubsets(nums []int, begin int, path []int, res *[][]int) {
	//1.满足条件

	r := make([]int, len(path))
	copy(r, path)
	fmt.Println(path)
	*res = append(*res, r)

	for i := begin; i < len(nums); i++ {
		// 选择i
		if len(path) > 0 && nums[i] <= path[len(path)-1] {
			//不能重复回头
			continue
		}
		path = append(path, nums[i])
		fmt.Println("递归前", path, i)
		backtraceSubsets(nums, begin+1, path, res)
		path = path[:len(path)-1]
		fmt.Println("递归后", path, i)
	}
}

//回溯
func backtraceSubsets2(nums []int, begin int, path []int, res *[][]int) {
	//1.满足条件

	r := make([]int, len(path))
	copy(r, path)
	fmt.Println(path)
	*res = append(*res, r)

	for i := begin; i < len(nums); i++ {
		path = append(path, nums[i])
		fmt.Println("递归前", path, i)
		newBegin := i + 1 //不能重复回头,和方法一达到一样的效果
		backtraceSubsets(nums, newBegin, path, res)
		path = path[:len(path)-1]
		fmt.Println("递归后", path, i)
	}
}
func subsets(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	sort.Ints(nums)
	backtraceSubsets(nums, 0, path, &res)
	return res
}
