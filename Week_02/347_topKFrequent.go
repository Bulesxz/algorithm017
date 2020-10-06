package main

import (
	"container/heap"
	"fmt"
)

/*
https://leetcode-cn.com/problems/top-k-frequent-elements/
347. 前 K 个高频元素
*/

type Item struct {
	Num   int
	Times int
}
type ItemMaxHeap []Item

func (h ItemMaxHeap) Len() int           { return len(h) }
func (h ItemMaxHeap) Less(i, j int) bool { return h[i].Times > h[j].Times }
func (h ItemMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ItemMaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Item))
}

func (h *ItemMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	fre := map[int]int{}
	for i := 0; i < len(nums); i++ {
		fre[nums[i]]++
	}

	h := &ItemMaxHeap{}
	heap.Init(h)
	for k, v := range fre {
		heap.Push(h, Item{k, v})
	}
	fmt.Println(h)

	result := []int{}
	for i := 0; i < k; i++ {
		t := heap.Pop(h)
		result = append(result, t.(Item).Num)
	}
	fmt.Println(result)
	return result
}
