package main

import "fmt"

func canReach(arr []int, start int) bool {
	visit := map[int]bool{}
	queue := []int{}
	queue = append(queue, start)
	visit[start] = true
	for len(queue) > 0 {
		curSize := len(queue)
		tmp := queue[:curSize]
		queue = []int{} // queue[:0] //tmp 和queue 分离 注意
		fmt.Println("queue:", queue, "tmp", tmp, "curSize", curSize)
		for j := 0; j < curSize; j++ {
			i := tmp[j]
			// fmt.Println("i", i, " arr[i]", arr[i])
			visit[i] = true
			if arr[i] == 0 {
				return true
			}
			right := i + arr[i]
			if right < len(arr) && !visit[right] {
				queue = append(queue, right)
			} else {
				fmt.Println("i", "right", i, right)
			}
			left := i - arr[i]
			if left >= 0 && !visit[left] {
				queue = append(queue, left)
			} else {
				fmt.Println("i", "left", i, left)
			}
			fmt.Println("queue2", queue)
		}
	}
	return false
}
