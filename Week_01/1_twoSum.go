package main

/*
1.两数之和
https://leetcode-cn.com/problems/two-sum/
*/

//使用map 存储,一次遍历 O(n)
func twoSum(nums []int, target int) []int {
	t := map[int]int{} //存储第二个目标值和下标  key:第二个目标值,value:下标
	for i := 0; i < len(nums); i++ {
		// fmt.Println(t, nums[i], target-nums[i])
		if index, ok := t[target-nums[i]]; ok { //目标值是否在map
			return []int{i, index}
		} else { //不在则加入map
			t[nums[i]] = i
		}
	}
	return nil
}

//两次遍历 O(n)
func twoSum2(nums []int, target int) []int {
	t := map[int]int{}
	for k, v := range nums {
		t[v] = k
	}
	for i := 0; i < len(nums); i++ {
		if index, ok := t[target-nums[i]]; ok { //目标值是否在map
			return []int{i, index}
		}
	}
	return nil
}

//暴力 两层循环 O{n2}
func twoSum3(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
