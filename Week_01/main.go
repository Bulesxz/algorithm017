package main

import "fmt"

func main() {
	nums := []int{2, 1, 11, 8, 15}
	target := 9
	result := twoSum3(nums, target)
	fmt.Println(result)

	digits := []int{1, 9, 9, 9}
	result = plusOne2(digits)
	fmt.Println(result)

	nums = []int{0, 1, 0, 3, 12}
	moveZeroes2(nums)
	fmt.Println(nums)
}
