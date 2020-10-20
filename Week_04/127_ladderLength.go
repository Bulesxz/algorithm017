package main

import (
	"fmt"
)

/*
https://leetcode-cn.com/problems/word-ladder/
127. 单词接龙
*/

//是否能够转换一个单词
func isChange(word string, wordList []string) []string {
	result := []string{}
	for i := 0; i < len(word); i++ {
		for j := 0; j < 26; j++ { //a-z
			c := []byte(word) // 将字符串 s 转换为 []byte 类型
			if c[i] != byte(j)+'a' {
				c[i] = byte(j) + 'a' //改变一份字符
			} else {
				continue
			}
			for k := range wordList {
				if string(c) == wordList[k] {
					result = append(result, string(c))
				}
			}
		}
	}
	return result
}

func genChangeA(L int, wordList []string) []string {
	result := []string{}
	for i := 0; i < len(wordList); i++ {
		for j := 0; j < L; j++ { //a-z
			c := []byte(wordList[i]) // 将字符串 s 转换为 []byte 类型
			c[j] = '*'
			result = append(result, string(c))
		}
	}
	return result
}

//BFS  无权无向图层级就是最短路径
func ladderLength(beginWord string, endWord string, wordList []string) int {
	visited := map[string]bool{}
	level := 0
	queue := []string{beginWord}
	visited[beginWord] = true
	for len(queue) > 0 {
		//遍历队列
		level++
		curSize := len(queue)
		for i := 0; i < curSize; i++ {
			top := queue[0]
			queue = queue[1:] //出栈
			next := isChange(top, wordList)
			for k := range next {
				if _, ok := visited[next[k]]; ok {
					//访问过
					continue
				}
				if endWord == next[k] {
					return level + 1
				}
				visited[next[k]] = true
				queue = append(queue, next[k])
				fmt.Println(queue)
			}
		}
	}
	return 0
}
func isChange2(tops, wordList []string) []string {
	result := []string{}
	for i := 0; i < len(tops); i++ {
		for j := 0; j < len(wordList); j++ {
			if tops[i] == wordList[j] {
				result = append(result, tops[i])
			}
		}
	}
	return result
}
func ladderLength2(beginWord string, endWord string, wordList []string) int {
	wordList = genChangeA(len(beginWord), wordList)
	fmt.Println(wordList)
	visited := map[string]bool{}
	level := 0
	queue := []string{beginWord}
	visited[beginWord] = true
	for len(queue) > 0 {
		//遍历队列
		level++
		curSize := len(queue)
		for i := 0; i < curSize; i++ {
			top := queue[0]
			queue = queue[1:] //出栈
			for k := 0; k < len(top); k++ {
				c := []byte(top) // 将字符串 s 转换为 []byte 类型
				c[k] = '*'
				newWord := string(c)
				fmt.Println("newWord", newWord, "old", top)
				has := false
				for m := 0; m < len(wordList); m++ {
					if wordList[m] == newWord {
						has = true
					}
				}
				if has == false {
					continue
				}
				if _, ok := visited[newWord]; ok {
					//访问过
					continue
				}
				if endWord == newWord {
					return level + 1
				}
				visited[newWord] = true
				queue = append(queue, newWord)
				fmt.Println(queue)
			}
		}
	}
	return 0
}
