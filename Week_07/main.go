package main

import "fmt"

func main() {
	// n := 3
	// res := solveNQueens(n)
	// fmt.Println(res)

	// start := "AACCGGTT"
	// end := "AACCGGTA"
	// bank := []string{"AACCGGTA"}

	start := "AACCTTGG"
	end := "AATTCCGG"
	bank := []string{"AATTCCGG", "AACCTGGG", "AACCCCGG", "AACCTACC"}
	res := minMutation2(start, end, bank)
	fmt.Println(res)
}
