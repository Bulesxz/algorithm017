package main

/*
46. 全排列
https://leetcode-cn.com/problems/permutations/
*/

// func helper(nums []int, res []int, used map[int]bool) {
// 	if len(res) == len(nums) {
// 		return
// 	}
// 	for i := 0; i < len(nums); i++ {
// 		if _, ok := used[nums[i]]; !ok {
// 			res = append(res, nums[i])
// 			used[nums[i]] = true
// 			fmt.Println(res)
// 			helper(nums, res, used)
// 			break
// 		}
// 	}
// }

func helper(nums []int, begin int, res *[][]int) {
	if begin == len(nums) {
		r := make([]int, len(nums))
		copy(r, nums)
		*res = append(*res, r)
	}
	for i := begin; i < len(nums); i++ {
		nums[begin], nums[i] = nums[i], nums[begin]
		helper(nums, begin+1, res)
		nums[begin], nums[i] = nums[i], nums[begin]
	}
}

//递归
func permute(nums []int) [][]int {
	res := [][]int{}
	helper(nums, 0, &res)
	return res
}
