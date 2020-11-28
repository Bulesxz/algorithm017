package main

/*
191. 位1的个数

求一个数末尾的1：n&(n-1)
*/
func hammingWeight(num uint32) int {
	cnt := 0
	for num != 0 {
		num = num & (num - 1)
		cnt++
	}
	return cnt
}
