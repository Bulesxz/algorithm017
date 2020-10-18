package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3, 4}
	// res := permute(nums)
	// nums := []int{1, 1, 2}
	// res := permuteUnique(nums)

	// res := generateParenthesis(3)
	// fmt.Println(res)

	// res := myPow(-2, 4)
	// nums := []int{1, 3}
	// res := subsets(nums)
	// res := subsetsWithDup(nums)
	digits := "23"
	res := letterCombinations2(digits)
	fmt.Println(res)

}
