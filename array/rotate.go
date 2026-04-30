package array

import (
	"slices"
)

//189. 轮转数组
//给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
//链接：https://leetcode.cn/problems/rotate-array/description/?envType=study-plan-v2&envId=top-100-liked
func rotate(nums []int, k int) {
	k = k % len(nums)
	slices.Reverse(nums)
	slices.Reverse(nums[:k]) //前k个轮转
	slices.Reverse(nums[k:]) //后n-k个轮转
}

func rotate2(nums []int, k int) {
	n := len(nums)
	k %= n
	cy := make([]int, n)
	copy(cy, nums)
	cy = append(cy, cy[:n-k]...)
	cy = cy[n-k:]
	copy(nums, cy)
}
