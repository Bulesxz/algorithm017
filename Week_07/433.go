package main

/*
433. 最小基因变化
https://leetcode-cn.com/problems/minimum-genetic-mutation/

*/

func inStrArry(i string, arr []string) bool {
	for item := range arr {
		if arr[item] == i {
			return true
		}
	}
	return false
}

//单向 BFS
func minMutation(start string, end string, bank []string) int {
	letter := []byte{'A', 'C', 'G', 'T'}
	q := []string{}
	q = append(q, start)
	visit := map[string]bool{}
	level := 0
	for len(q) > 0 {
		fronts := q[:]
		q = []string{}
		for _, front := range fronts {
			if front == end {
				return level
			}
			visit[front] = true
			for i := 0; i < len(front); i++ {
				tmp := []byte(front)
				for j := range letter {
					if letter[j] != front[i] {
						tmp[i] = letter[j]
						if !visit[string(tmp)] && inStrArry(string(tmp), bank) {
							q = append(q, string(tmp))
						}
					}
				}
			}
		}
		level++
	}
	return -1
}

//双向BFS
func minMutation2(start string, end string, bank []string) int {
	letter := []byte{'A', 'C', 'G', 'T'}
	q := []string{} //开始
	q = append(q, start)
	q2 := []string{} //结束
	q2 = append(q2, end)
	visit := map[string]bool{}
	level := 0
	for len(q) > 0 {
		fronts := q[:]
		q = []string{}
		for _, front := range fronts {

			if inStrArry(front, q2) {
				return level
			}
			visit[front] = true
			for i := 0; i < len(front); i++ {
				tmp := []byte(front)
				for j := range letter {
					if letter[j] != front[i] {
						tmp[i] = letter[j]
						if !visit[string(tmp)] && inStrArry(string(tmp), bank) {
							q = append(q, string(tmp))
						}
					}
				}
			}
		}
		level++
		if len(q) > len(q2) {
			q, q2 = q2, q
		}
	}
	return -1
}
