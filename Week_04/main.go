package main

import (
	"fmt"
)

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	res := ladderLength2(beginWord, endWord, wordList)
	fmt.Println(res)

	// res1 := genChangeA(len(beginWord), wordList)
	// fmt.Println(res1)
}
