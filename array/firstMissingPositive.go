package array

//41. 缺失的第一个正数
//给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
//请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
//链接：https://leetcode.cn/problems/first-missing-positive/description/?envType=study-plan-v2&envId=top-100-liked
func firstMissingPositive(nums []int) int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for nums[i] > 0 && nums[i] <= length && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	for i := 0; i < length; i++ {
		if nums[i] != nums[i+1] {
			return i + 1
		}
	}

	return length + 1
}
