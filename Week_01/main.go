package main

import "fmt"

func main() {
	nums := []int{2, 1, 11, 8, 15}
	target := 9
	result := twoSum3(nums, target)
	fmt.Println(result)
}
