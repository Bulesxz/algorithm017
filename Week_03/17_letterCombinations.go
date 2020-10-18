package main

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
	res := []string{}

	queue := []string{}
	queue = append(queue, numbers[string(digits[0])]...)
	index := 1

	for len(queue) > 0 {
		nums := queue[0]
		queue = queue[1:]
		//扩展子节点
		nexts := numbers[string(digits[index])]
		for _, a := range nexts {
			queue = append(queue, nums+a)
		}
	}

	return res
}
