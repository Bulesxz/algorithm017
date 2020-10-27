package main

/*
322. 零钱兑换
https://leetcode-cn.com/problems/coin-change/
*/

/*
方法一
递归: 每一个硬币选或者必选 爬楼梯
*/
func recursiveCoinChange() {

}
func coinChange(coins []int, amount int) int {
	//1. 递归终止
	if amount < 0 {
		return
	}
	for coin := range coins {
		a := coinChange(coins, amount-coin)
		b := coinChange(coins, amount)
		if a < b {
			max = a
		} else {
			max = b
		}
	}
}
