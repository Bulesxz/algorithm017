package main

/*
50. Pow(x, n)
https://leetcode-cn.com/problems/powx-n/
*/

//分治法

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n%2 == 0 {
		num := quickMul(x, n/2)
		return num * num
	} else {
		num := quickMul(x, n/2)
		return num * num * x
	}
}
func myPow(x float64, n int) float64 {
	flagn := 0
	if n < 0 {
		flagn = 1
		n *= -1
	}
	num := quickMul(x, n)
	if flagn == 1 {
		num = 1 / num
	}

	return num
}
