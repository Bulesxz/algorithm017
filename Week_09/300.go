package main

import (
	"fmt"
)

/*
300. 最长上升子序列
dp 为 0到第i个元素的最长子序列
dp[i] = max(dp[j]) + 1

o(N2)
*/
func lengthOfLIS(nums []int) int {
	if len(nums) < 1 {
		return len(nums)
	}
	//
	dp := make([]int, len(nums))
	dp[0] = 1
	MAX := 1
	for i := 1; i < len(nums); i++ {
		max := dp[0]
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > max {
					max = dp[j] + 1
				}
			}
		}
		dp[i] = max
		if max > MAX {
			MAX = max
		}
		fmt.Println(i, dp)
	}
	return MAX
}

/*
贪心算法
dp[i] 为第i+1 长度的子序列  最小的nums元素
*/
func lengthOfLIS2(nums []int) int {
	if len(nums) < 1 {
		return len(nums)
	}
	//
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	end := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > dp[end] { //找到了更长的子序列
			end++
			dp[end] = nums[i]
		} else {
			//调整之前的最小值
			rignt := end
			left := 0
			for left < rignt {
				mid := left + (rignt-left)/2
				if dp[mid] < nums[i] {
					left = mid + 1
				} else {
					rignt = mid
				}
			}
			dp[left] = nums[i]
		}
	}
	return end + 1
}
