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

func genChangeA(L int, wordList []string) map[string]map[string]bool {
	result := map[string]map[string]bool{}
	for i := 0; i < len(wordList); i++ {
		for j := 0; j < L; j++ { //a-z
			c := []byte(wordList[i]) // 将字符串 s 转换为 []byte 类型
			c[j] = '*'
			if _, ok := result[string(c)]; ok {
				result[string(c)][wordList[i]] = true
			} else {
				result[string(c)] = map[string]bool{wordList[i]: true}
			}
		}
	}
	return result
}

func ladderLength2(beginWord string, endWord string, wordList []string) int {
	wordmap := genChangeA(len(beginWord), wordList)
	// fmt.Println(wordmap)
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

			next := map[string]bool{}
			for k := 0; k < len(top); k++ {
				c := []byte(top) // 将字符串 s 转换为 []byte 类型
				c[k] = '*'
				newWord := string(c)
				if t, ok := wordmap[newWord]; ok {
					for key := range t {
						next[key] = true
					}
				}
			}
			// fmt.Println("next", next)

			for w := range next {
				if _, ok := visited[w]; ok {
					//访问过
					continue
				}
				if endWord == w {
					return level + 1
				}
				visited[w] = true
				queue = append(queue, w)
				// fmt.Println("queue", queue)
			}

		}
	}
	return 0
}
