package main

/*
126. 单词接龙 II
https://leetcode-cn.com/problems/word-ladder-ii/description/
*/

// func findLadders(beginWord string, endWord string, wordList []string) [][]string {
// 	wordmap := genChangeA(len(beginWord), wordList)
// 	// fmt.Println(wordmap)
// 	visited := map[string]bool{}
// 	level := 0
// 	queue := []string{beginWord}
// 	visited[beginWord] = true
// 	for len(queue) > 0 {
// 		//遍历队列
// 		level++
// 		curSize := len(queue)
// 		for i := 0; i < curSize; i++ {
// 			top := queue[0]
// 			queue = queue[1:] //出栈

// 			next := map[string]bool{}
// 			for k := 0; k < len(top); k++ {
// 				c := []byte(top) // 将字符串 s 转换为 []byte 类型
// 				c[k] = '*'
// 				newWord := string(c)
// 				if t, ok := wordmap[newWord]; ok {
// 					for key := range t {
// 						next[key] = true
// 					}
// 				}
// 			}
// 			// fmt.Println("next", next)

// 			for w := range next {
// 				if _, ok := visited[w]; ok {
// 					//访问过
// 					continue
// 				}
// 				if endWord == w {
// 					return level + 1
// 				}
// 				visited[w] = true
// 				queue = append(queue, w)
// 				// fmt.Println("queue", queue)
// 			}

// 		}
// 	}
// 	return 0
// }
