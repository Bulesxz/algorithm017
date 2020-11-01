package main

import (
	"fmt"
)

func main() {
	// beginWord := "hit"
	// endWord := "cog"
	// wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	// res := ladderLength2(beginWord, endWord, wordList)
	// fmt.Println(res)

	// res1 := genChangeA(len(beginWord), wordList)
	// fmt.Println(res1)

	// intervals := [][]int{[]int{2, 6}, []int{1, 3}, []int{8, 10}, []int{15, 18}}
	// intervals := [][]int{[]int{1, 4}, []int{2, 3}}
	intervals := [][]int{[]int{1, 4}, []int{0, 2}, []int{3, 5}}
	res2 := merge(intervals)
	fmt.Println(res2)
}
