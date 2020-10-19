package main

import (
	"fmt"
)

func backtraceLetterCombinations(depth int, digits string, path string, numbers map[string][]string, res *[]string) {
	//1.满足条件
	if depth == len(digits) {
		*res = append(*res, path)
		return
	}
	//2.选择路径
	for _, a := range numbers[string(digits[depth])] {
		//利用传path+a 传局部变量 省去撤销步骤书写
		backtraceLetterCombinations(depth+1, digits, path+a, numbers, res)
	}
}

//回溯法
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	numbers := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	res := []string{}

	backtraceLetterCombinations(0, digits, "", numbers, &res)
	return res
}

//广度优先
func letterCombinations2(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	numbers := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}

	queue := []string{}
	queue = append(queue, "")
	for j := 0; j < len(digits); j++ { //bfs 层数
		cnt := len(queue) //当前层队列长度
		fmt.Println("====", cnt, string(digits[j]))
		for i := 0; i < cnt; i++ { //当前层队列元素
			fmt.Println("nums", queue, len(queue))
			nums := queue[0]  //队列头
			queue = queue[1:] //出队列
			// fmt.Println("nums", nums, queue, len(queue))
			//扩展子节点
			nexts := numbers[string(digits[j])]
			for _, a := range nexts {
				queue = append(queue, nums+a)
			}
		}
	}

	return queue
}
