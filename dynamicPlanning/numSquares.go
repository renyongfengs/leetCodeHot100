package dynamicPlanning

import (
	"math"
)

//279. 完全平方数
//给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
//完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
//链接：https://leetcode.cn/problems/perfect-squares/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/perfect-squares/solutions/823192/dai-ma-sui-xiang-lu-279-wan-quan-ping-fa-9ieo/?envType=study-plan-v2&envId=top-100-liked
func numSquares(n int) int {
	result := make([]int, n+1)
	for i := 1; i <= n; i++ {
		result[i] = math.MaxInt
	}
	for i := 1; i <= n; i++ {
		for j := i * i; j <= n; j++ {
			result[j] = min(result[j], result[j-i*i]+1)
		}
	}
	return result[n]
}
