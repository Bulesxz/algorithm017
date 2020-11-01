package main

import (
	"fmt"
	"sort"
)

/*
56. 合并区间
https://leetcode-cn.com/problems/merge-intervals/

示例 1:

输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

*/

type Arry [][]int

//Len()
func (s Arry) Len() int {
	return len(s)
}

//Less(): 首个元素
func (s Arry) Less(i, j int) bool {
	if s[i][0] < s[j][0] {
		return true
	} else if s[i][0] == s[j][0] {
		if s[i][len(s[i])-1] <= s[j][len(s[j])-1] {
			return true
		}
	}
	return false
}

//Swap()
func (s Arry) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

/*
1. 整个数组排序，顺序安装子数组首个元素
O(nlogn)
*/
func merge(intervals [][]int) [][]int {
	sort.Sort(Arry(intervals))
	fmt.Println(intervals)
	for i := 0; i < len(intervals)-1; i++ {

		//还有下一个区间
		min := intervals[i][0] //本区间最大元素
		max := intervals[i][1] //本区间最大元素
		fmt.Println("max", max)
		nextMin := intervals[i+1][0]
		nextMax := intervals[i+1][1]
		// fmt.Println("max", max, "nextMin", nextMin, "nextMax", nextMax)
		if (max >= nextMin && max <= nextMax) || (nextMax >= min && nextMax <= max) {
			if intervals[i+1][1] > intervals[i][1] {
				intervals[i][1] = intervals[i+1][1] //合并
			}
			intervals = append(intervals[:i+1], intervals[i+2:]...) //删除第i个
			// fmt.Println("========", intervals)
			i--
		}

	}
	return intervals
}
