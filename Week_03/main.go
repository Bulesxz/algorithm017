package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3, 4}
	// res := permute(nums)
	nums := []int{1, 1, 2}
	res := permuteUnique(nums)
	fmt.Println(res)
}
