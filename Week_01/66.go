package main

/*
66 加一
https://leetcode-cn.com/problems/plus-one/
*/

// 遍历每次都判断进位
//O(n)
func plusOne(digits []int) []int {
	c := 1 //余数，这里为初始化的加一
	for i := len(digits) - 1; i >= 0; i-- {
		c1 := (digits[i] + c) / 10       //进位
		digits[i] = (digits[i] + c) % 10 //取整
		c = c1
		if c == 0 { //没有进位则直接返回
			return digits
		}
	}
	if c != 0 { //进位
		digits = append([]int{c}, digits...)
	}
	return digits
}

//O(n)
//O(n)
func plusOne2(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			//不会存在进位
			return digits
		}
		digits[i] = 0
		//继续遍历
	}

	//之前没有return，说明数组长度要+1
	digits = append([]int{1}, digits...)
	return digits
}
