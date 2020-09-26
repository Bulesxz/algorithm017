package main

/*
11. 盛最多水的容器
https://leetcode-cn.com/problems/container-with-most-water/
*/

//双指针
func maxArea(height []int) int {
	minHeight := func(i, j int) int {
		if height[j] < height[i] {
			return height[j]
		} else {
			return height[i]
		}
	}
	max := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			area := (j - i) * minHeight(i, j)
			if area > max {
				max = area
			}
		}
	}
	return max
}

func maxArea2(height []int) int {
	minHeight := func(i, j int) int {
		if height[j] < height[i] {
			return height[j]
		} else {
			return height[i]
		}
	}
	max := 0
	for i, j := 0, len(height)-1; i < j; {
		area := (j - i) * minHeight(i, j)
		if area > max {
			max = area
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return max
}
