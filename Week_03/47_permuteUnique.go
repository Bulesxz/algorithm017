package main

import "sort"

/*
47. 全排列 II
https://leetcode-cn.com/problems/permutations-ii/
*/
func backtracePermuteUnique(nums []int, used []bool, path []int, dep int, res *[][]int) {
	//1. 终止条件
	if dep == len(nums) {
		r := make([]int, len(nums))
		copy(r, path)
		// fmt.Println(r)
		*res = append(*res, r)
		return
	}
	//2.路径选择
	for i := 0; i < len(nums); i++ {
		if used[i] {
			//被使用过
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			//同一层级已经被使用过
			continue
		}
		used[i] = true //进入
		path = append(path, nums[i])
		// fmt.Println("进入递归前:", used, i)
		backtracePermuteUnique(nums, used, path, dep+1, res)
		used[i] = false //撤销
		path = path[:len(path)-1]
		// fmt.Println("进入递归后:", used, i)
	}
}

//回溯
func permuteUnique(nums []int) [][]int {
	used := make([]bool, len(nums))
	res := [][]int{}
	path := []int{}
	sort.Ints(nums) //排序
	backtracePermuteUnique(nums, used, path, 0, &res)
	return res
}
