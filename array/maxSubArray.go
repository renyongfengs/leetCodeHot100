package array

import (
	"math"
)

//53. 最大子数组和
//给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//子数组是数组中的一个连续部分。
//链接：https://leetcode.cn/problems/maximum-subarray/description/?envType=study-plan-v2&envId=top-100-liked
//讲解：https://leetcode.cn/problems/maximum-subarray/?envType=study-plan-v2&envId=top-100-liked
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := math.MinInt
	preSum, minPreSum := 0, 0
	for _, val := range nums {
		preSum += val                          //前缀和
		result = max(result, preSum-minPreSum) //当前位置的最大子数组长度和当前结果比较大小

		minPreSum = min(minPreSum, preSum) //维护前缀和的最小值
	}
	return result
}
