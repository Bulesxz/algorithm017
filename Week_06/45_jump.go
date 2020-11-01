package main

import (
	"fmt"
	"math"
)

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// //1.从begin 跳到结束 要跳最少的次数
// func recJump(nums []int, begin int, cache map[int]int) int {
// 	if begin >= len(nums)-1 {
// 		return 0
// 	}
// 	if min, ok := cache[begin]; ok {
// 		fmt.Println("fit")
// 		return min
// 	}
// 	fmt.Println("begin", begin, "step", nums[begin])
// 	min := math.MaxInt32
// 	for i := 1; i <= nums[begin]; i++ {
// 		step := recJump(nums, begin+i, cache) + 1
// 		if step < min {
// 			min = step
// 		}
// 	}
// 	cache[begin] = min
// 	return min
// }
// func jump(nums []int) int {
// 	cache := map[int]int{}
// 	return recJump(nums, 0, cache)
// }

/*
dp[i] = min (dp[i-1],dp[i-2],..0) +1
*/
func jump(nums []int) int {
	dp := make([]int, len(nums), len(nums)) //从 i到末尾最小步数
	for i := 0; i < len(nums); i++ {
		dp[i] = len(nums) - i
	}
	dp[0] = 0
	for i := 1; i < len(nums); i++ {
		min := math.MaxInt32
		for j := 0; j <= i-1; j++ {
			fmt.Println("j", j, "i", i)
			if i-j <= nums[j] && dp[j] < min {
				//从j 一步到i
				//长度小于 ,一步到位
				min = dp[j]
			}
			fmt.Println("j", j, ">>>i", i, dp[i])
		}
		dp[i] = min + 1
	}
	return dp[len(nums)-1]
}
