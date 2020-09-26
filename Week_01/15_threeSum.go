package main

import "fmt"

/*
15. 三数之和
https://leetcode-cn.com/problems/3sum/
*/

//暴力

func twoSums(offset int, nums []int, target int) [][]int {
	t := map[int]int{}
	result := [][]int{}
	for i := offset; i < len(nums); i++ {
		if index, ok := t[target-nums[i]]; ok { //目标值是否在map

			result = append(result, []int{i, index})

		} else {
			t[nums[i]] = i
		}
	}
	return result
}

// [-1,0,1,2,-1,-4]
func threeSum(nums []int) [][]int {
	result := map[string][]int{}
	for i := 0; i < len(nums)-2; i++ {
		r := twoSums(i+1, nums, 0-nums[i])
		for k := range r {
			tmp := []int{nums[r[k][0]], nums[r[k][1]], nums[i]}
			quickSort(tmp, 0, len(tmp)-1)
			key := fmt.Sprintf("%d_%d_%d", tmp[0], tmp[1], tmp[2]) //需要去重
			result[key] = tmp
		}
	}
	result2 := [][]int{}
	for k := range result {
		result2 = append(result2, result[k])
	}
	return result2
}

func threeSum2(nums []int) [][]int {
	result := [][]int{}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println("====", nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			for j < k {
				if nums[i]+nums[j]+nums[k] > 0 {
					k--
				} else if nums[i]+nums[j]+nums[k] < 0 {
					j++
				} else {
					result = append(result, []int{nums[i], nums[j], nums[k]})
					j++
					k--
					for j < k && nums[j] == nums[j-1] {
						j++
					}
					for j < k && nums[k] == nums[k+1] {
						k--
					}
				}
			}
		}

	}
	return result
}
