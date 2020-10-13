package main

import "fmt"

/*
46. 全排列
https://leetcode-cn.com/problems/permutations/
*/

func helper(nums []int, res []int, used map[int]bool) {
	if len(res) == len(nums) {
		return
	}
	for i := 0; i < len(nums); i++ {
		if _, ok := used[nums[i]]; !ok {
			res = append(res, nums[i])
			used[nums[i]] = true
			fmt.Println(res)
			helper(nums, res, used)
			break
		}
	}
}
func permute(nums []int) [][]int {
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		r := []int{nums[i]}
		helper(nums, r, map[int]bool{nums[i]: true})
		res = append(res, r)
	}
	return res
}
