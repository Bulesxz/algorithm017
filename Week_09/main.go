package main

import "fmt"

func main() {
	// s := "abca"
	// res := validPalindrome(s)
	// fmt.Println(res)

	nums := []int{4, 10, 4, 3, 8, 9}
	res := lengthOfLIS2(nums)
	fmt.Println(res)
}
