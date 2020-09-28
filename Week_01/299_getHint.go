package main

import "fmt"

/*
	299. 猜数字游戏
	https://leetcode-cn.com/problems/bulls-and-cows/
*/
func getHint(secret string, guess string) string {
	guessA := 0
	minLen := len(secret)
	if len(secret) > len(guess) {
		minLen = len(guess)
	}

	secretMap := map[byte]int{}
	for i := 0; i < len(secret); i++ {
		secretMap[secret[i]]++
	}
	guessMap := map[byte]int{}
	for i := 0; i < len(guess); i++ {
		guessMap[guess[i]]++
	}

	// secret := "1123"
	// guess := "0111"

	Amap := map[byte]int{}
	for i := 0; i < minLen; i++ {
		if secret[i] == guess[i] {
			guessA++
			Amap[secret[i]] = Amap[secret[i]] + 1
		}
	}

	fmt.Println(secretMap, guessMap)

	guessB := 0
	for k, v := range secretMap {
		if g, ok := guessMap[k]; ok {

			if g < v {
				guessB = guessB + g
			} else {
				guessB = guessB + v
			}
			if a, ok := Amap[k]; ok {
				guessB = guessB - a
			}
		}
	}
	return fmt.Sprintf("%dA%dB", guessA, guessB)
}

func getHint2(secret string, guess string) string {
	guessA := 0
	guessMap := map[byte]int{}
	// secret := "1123"
	// guess := "0111"
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			guessA++ //公牛
		} else {
			guessMap[secret[i]]++
			guessMap[guess[i]]--
		}
	}

	noguess := 0
	for _, v := range guessMap {
		if v > 0 { //没有猜对的数字
			noguess = noguess + v
		}
	}
	guessB := len(secret) - noguess
	return fmt.Sprintf("%dA%dB", guessA, guessB)
}
