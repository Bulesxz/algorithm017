package main

import (
	"fmt"
	"sort"
)

/*
90. 子集 II
https://leetcode-cn.com/problems/subsets-ii/
*/

func backtraceSubsetsWithDup(nums []int, begin int, path []int, res *[][]int) {
	//1.满足条件
	r := make([]int, len(path))
	copy(r, path)
	// fmt.Println(path)
	*res = append(*res, r) //都满足
	for i := begin; i < len(nums); i++ {

		if i > begin && nums[i] == nums[i-1] {
			// i>beigin 说明是回溯到第二个元素了
			continue
		}
		fmt.Println("递归前", path, nums[i], "进入")
		path = append(path, nums[i])
		// new_begin:= begin+1 错误，这样导致选第二个元素时候，第二个元素已经被第一层使用了
		new_begin := i + 1 //下一层永远要从本层后面的元素开始，这样都是第一层之后的元素
		backtraceSubsetsWithDup(nums, new_begin, path, res)
		path = path[:len(path)-1]
		fmt.Println("递归后", path, nums[i], "出来")
	}
}
func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	sort.Ints(nums)
	backtraceSubsetsWithDup(nums, 0, path, &res)
	return res
}
