package main

import (
	"math"
)

/*
322. 零钱兑换
https://leetcode-cn.com/problems/coin-change/
*/

/*
方法一
递归: 每一个硬币选或者必选 类似爬楼梯
*/
func recursiveCoinChange(coins []int, amount int, memo map[int]int) int {
	if res, ok := memo[amount]; ok {
		return res
	}
	//1. 递归终止
	if amount < 0 {
		memo[amount] = -1
		return -1
	}
	if amount == 0 {
		memo[amount] = 0
		return 0
	}
	min := math.MaxInt32
	for _, coin := range coins {
		// fmt.Println(amount)
		res := recursiveCoinChange(coins, amount-coin, memo)
		// fmt.Println("amount", amount, "min:", min, coin)
		if res >= 0 && res < min {
			min = res + 1 // 加1，是为了加上得到res结果的那个步骤中，兑换的一个硬币
		}
	}

	if min == math.MaxInt32 {
		memo[amount] = -1
		return -1
	} else {
		memo[amount] = min
		return min
	}
}
func coinChange(coins []int, amount int) int {
	memo := map[int]int{}
	return recursiveCoinChange(coins, amount, memo)
}
