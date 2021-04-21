package dynamic_programming

// LeetCode第70题：https://leetcode-cn.com/problems/climbing-stairs/

func ClimbStairs(n int) int {
	tempMap := make(map[int]int)
	tempMap[1] = 1
	tempMap[2] = 2
	for i := 3; i <= n; i++ {
		tempMap[i] = tempMap[i-1] + tempMap[i-2]
	}
	return tempMap[n]
}
