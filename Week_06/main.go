package main

import "fmt"

func main() {
	m := 3
	n := 2
	res := uniquePaths2(m, n)
	fmt.Println(res)
	coins := []int{1, 2, 5}
	res2 := coinChange(coins, 100)
	fmt.Println(res2)
}
