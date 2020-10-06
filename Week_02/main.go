package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 4, 5, 1, 3, -1, 2, 3, 0}
	fmt.Println("=============", nums)

	// h := &IntHeap{100, 16, 4, 8, 70, 2, 36, 22, 5, 12}

	// fmt.Println("\nHeap:")
	// heap.Init(h)
	// heap.Push(h, 1)
	// heap.Push(h, 1)
	// heap.Push(h, 1)
	// heap.Push(h, 1)
	// heap.Push(h, 1)
	// heap.Push(h, 1)
	// fmt.Println(heap.Pop(h))
	// fmt.Println(heap.Pop(h))
	// fmt.Println(heap.Pop(h))
	// fmt.Println(heap.Pop(h))
	// fmt.Println(heap.Pop(h))

	// fmt.Printf("最大值: %d\n", (*h)[0])

	ugly := nthUglyNumber(10)
	fmt.Println(ugly)

	nums = []int{1, 1, 1, 2, 2, 3}
	k := 2
	topKFrequent(nums, k)

}
