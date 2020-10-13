package main

import (
	"container/heap"
	"fmt"
)

/*
剑指 Offer 49. 丑数

https://leetcode-cn.com/problems/chou-shu-lcof/
*/

//使用小顶堆
func nthUglyNumber(n int) int {

	h := &IntHeap{1}
	vist := map[interface{}]struct{}{}
	cnt := 0
	heap.Init(h)
	for len(vist) < n {
		min := heap.Pop(h)
		fmt.Println("==========数字:", min)
		if _, ok := vist[min]; !ok {
			heap.Push(h, min.(int)*2)
			heap.Push(h, min.(int)*3)
			heap.Push(h, min.(int)*5)
			fmt.Println(ok, "arr:", *h)
			vist[min] = struct{}{}
			cnt++
			if cnt == n {
				return min.(int)
			}
		}
	}
	return 0
}

//迭代
func nthUglyNumber2(n int) int {
	p2, p3, p5 := 0, 0, 0
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp2 := dp[p2*2]
		dp3 := dp[p3*2]
		dp5 := dp[p5*2]
		min := dp2
		if dp3 < min {
			min = dp3
		}
		if dp5 < min {
			min = dp5
		}
		dp[i] = min
		if dp[i] == dp2 {
			p2++
		}
		if dp[i] == dp3 {
			p3++
		}
		if dp[i] == dp5 {
			p5++
		}
	}
	return dp[n-1]
}
