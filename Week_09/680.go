package main

import (
	"fmt"
)

/*
680. 验证回文字符串 Ⅱ
*/
func validPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	cnt := 0
	for i <= j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			cnt++
			//删除i 或者删除j
			tmpi := i
			tmpj := j
			flag1 := true
			for tmpi++; tmpi <= tmpj; {
				fmt.Println(tmpi, tmpi, s[tmpi], s[tmpj])
				if s[tmpi] == s[tmpj] {
					tmpi++
					tmpj--
				} else {
					flag1 = false
					break
				}
			}
			if flag1 == true {
				return true
			}
			tmpi = i
			tmpj = j
			flag2 := true
			for tmpj--; tmpi <= tmpj; {
				if s[tmpi] == s[tmpj] {
					tmpi++
					tmpj--
				} else {
					flag2 = false
					break
				}
			}
			if flag2 == true {
				return true
			}
			return flag1 || flag2
		}
	}
	return cnt <= 1
}
