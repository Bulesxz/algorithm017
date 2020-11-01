package main

import (
	"fmt"
	"sort"
)

/*
4数之和
*/

func threeSum(nums []int, target int) [][]int {
	fmt.Println("nums", nums)
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		s := target - nums[i]
		begin := i + 1
		end := len(nums) - 1
		// fmt.Println("begin", begin, "end", end)
		for begin < end {
			if nums[begin]+nums[end] == s {
				r := []int{}
				r = append(r, nums[i], nums[begin], nums[end])
				res = append(res, r)
				fmt.Println("res", res, nums[i], i)
				for begin < end && nums[begin] == nums[begin+1] {
					begin++
				}
				begin++

				for begin < end && nums[end] == nums[end-1] {
					end--
				}
				end--

				//ok
			} else if nums[begin]+nums[end] < s {
				begin++
				// for begin < end && nums[begin] == nums[begin-1] {
				// 	begin++
				// } 不需要，肯定不等于target
			} else {
				end--
				// for begin < end && nums[end] == nums[end+1] {
				// 	end--
				// } 不需要，肯定不等于target
			}
		}
	}
	return res
}

func fourSum(nums []int, target int) [][]int {
	//1.排序
	sort.Ints(nums)
	// fmt.Println(nums)
	var res [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		r := threeSum(nums[i+1:], target-nums[i])
		for j := 0; j < len(r); j++ {
			t := []int{nums[i]}
			t = append(t, r[j]...)
			res = append(res, t)
		}
	}
	return res
}
