package array

//238. 除自身以外数组的乘积
//给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
//题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
//请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
//链接 ：https://leetcode.cn/problems/product-of-array-except-self/?envType=study-plan-v2&envId=top-100-liked
func productExceptSelf(nums []int) []int {
	pre := make([]int, len(nums))
	pre[0] = 1
	for i := 1; i < len(nums); i++ {
		pre[i] = pre[i-1] * nums[i-1]
	}

	suf := make([]int, len(nums))
	suf[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		suf[i] = suf[i+1] * nums[i+1]
	}

	result := make([]int, len(nums))
	for idx, val := range suf {
		result[idx] = pre[idx] * val
	}

	return result
}

func productExceptSelf2(nums []int) []int {
	suf := make([]int, len(nums))
	suf[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		suf[i] = suf[i+1] * nums[i+1]
	}

	pre := 1
	for idx, val := range nums {
		suf[idx] *= pre
		pre *= val
	}

	return suf
}
