package subString

//560. 和为 K 的子数组
//给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
//子数组是数组中元素的连续非空序列。
//链接：https://leetcode.cn/problems/subarray-sum-equals-k/description/?envType=study-plan-v2&envId=top-100-liked

func subarraySum(nums []int, k int) int {
	result := 0

	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j >= 0; j-- {
			sum += nums[j]
			if sum == k {
				result++
			}
		}
	}
	return result
}

func subarraySum2(nums []int, k int) int {
	result := 0
	sum := 0
	vm := make(map[int]int)
	vm[0] = 1 //值本身就可以
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if count, ok := vm[sum-k]; ok {
			result += count
		}

		vm[sum]++
	}

	return result
}
