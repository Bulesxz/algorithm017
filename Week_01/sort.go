package main

func partation(nums []int, left, right int) int {
	key := nums[left]
	for left < right {
		for left < right && nums[right] >= key {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] < key {
			left++
		}
		nums[right] = nums[left]

	}
	nums[left] = key
	return left
}

//快速排序
func quickSort(nums []int, beign, end int) {
	if beign < end {
		middle := partation(nums, beign, end)
		quickSort(nums, beign, middle-1)
		quickSort(nums, middle+1, end)
	}
}
