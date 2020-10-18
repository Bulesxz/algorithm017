package main

import (
	"sort"
)

/*
169. 多数元素
https://leetcode-cn.com/problems/majority-element/description/
*/

//排序之后 中间位置为众数
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[(len(nums)-1)/2]
}
