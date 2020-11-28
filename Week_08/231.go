package main

/*
231. 2的幂
最高位为1 其他位置为0
*/
func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}
	return n&(n-1) == 0
}
