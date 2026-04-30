package dynamicPlanning

import (
	"math"
)

//152. 乘积最大子数组
//给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续 子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
//测试用例的答案是一个 32-位 整数。
//请注意，一个只包含一个元素的数组的乘积是这个元素的值
//链接：https://leetcode.cn/problems/maximum-product-subarray/?envType=study-plan-v2&envId=top-100-liked

func maxProduct(nums []int) int {
	n := len(nums)
	fMax := make([]int, n)
	fMin := make([]int, n)
	fMax[0], fMin[0] = nums[0], nums[0]
	ans := fMax[0]
	for i := 1; i < n; i++ {
		x := nums[i]
		// 把 x 加到右端点为 i-1 的（乘积最大/最小）子数组后面，
		// 或者单独组成一个子数组，只有 x 一个元素
		fMax[i] = max(fMax[i-1]*x, fMin[i-1]*x, x)
		fMin[i] = min(fMax[i-1]*x, fMin[i-1]*x, x)
		ans = max(ans, fMax[i])
	}
	return ans
}

func maxProduct2(nums []int) int {
	ans := math.MinInt
	fmax, fmin := 1, 1
	for _, num := range nums {
		fmax, fmin = max(num, fmax*num, fmin*num), min(num, fmax*num, fmin*num)
		ans = max(ans, fmax)
	}
	return ans
}
