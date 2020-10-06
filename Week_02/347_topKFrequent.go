package main

import (
	"container/heap"
)

/*
https://leetcode-cn.com/problems/top-k-frequent-elements/
347. 前 K 个高频元素
*/

type Item struct {
	Num   int
	Times int
}
type IteminHeap []Item

func (h IteminHeap) Len() int           { return len(h) }
func (h IteminHeap) Less(i, j int) bool { return h[i].Times < h[j].Times }
func (h IteminHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IteminHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Item))
}

func (h *IteminHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IteminHeap) Top() interface{} {
	// n := len(*h)
	x := (*h)[0]
	return x
}

func topKFrequent(nums []int, k int) []int {
	fre := map[int]int{}
	for i := 0; i < len(nums); i++ {
		fre[nums[i]]++
	}

	h := &IteminHeap{}
	heap.Init(h)
	cnt := 0
	for num, times := range fre {
		if cnt >= k {
			if h.Top().(Item).Times < times {
				heap.Pop(h)
			} else {
				continue
			}
		}
		heap.Push(h, Item{num, times})
		cnt++
	}

	result := []int{}
	for i := 0; i < k; i++ {
		t := heap.Pop(h)
		result = append(result, t.(Item).Num)
	}
	return result
}
