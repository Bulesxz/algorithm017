package main

/*
283 移动零
https://leetcode-cn.com/problems/move-zeroes/
*/

// O(n)
func moveZeroes(nums []int) {
	j := 0 // 下一个非0元素该放置的位置
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			//不为0,则放置到j
			nums[j] = nums[i]
			if j != i {
				//i 位置要置0
				nums[i] = 0
			}
			j++ //下一个非0该放的位置
		}
	}
}

//非0元素个数即下一个非0元素改移动到的位置
func moveZeroes2(nums []int) {
	notZeroCnt := 0 //非元素个数
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[notZeroCnt], nums[i] = nums[i], nums[notZeroCnt]
			notZeroCnt++
		}
	}
}
