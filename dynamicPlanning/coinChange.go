package dynamicPlanning

import (
	"math"
)

//322. 零钱兑换
//给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
//计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
//你可以认为每种硬币的数量是无限的。
//链接：https://leetcode.cn/problems/coin-change/description/?envType=study-plan-v2&envId=top-100-liked
func coinChange(coins []int, amount int) int {
	result := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		result[i] = math.MaxInt32
	}
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}

			result[i] = min(result[i-coin]+1, result[i])
		}
	}

	if result[amount] == math.MaxInt32 {
		return -1
	}

	return result[amount]
}
