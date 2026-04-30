package twoPointers

import "sort"

//15. 三数之和
//给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
//注意：答案中不可以包含重复的三元组。
//链接：https://leetcode.cn/problems/3sum/description/?envType=study-plan-v2&envId=top-100-liked

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) < 3 || nums == nil {
		return result
	}
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < len(nums)-2; i++ {
		//最小值大于0，不存在累加和小于0的可能
		if nums[i] > 0 {
			break
		}

		//去重第一层
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		//nums[i]和后面两个最小的值相加大于0，说明nums[i]太大，不存在累加和为0的可能
		if nums[i]+nums[i+1]+nums[i+2] > 0 {
			break
		}
		//nums[i]和最大的值相加小于0，说明nums[i]太小，不存在累加和为0的可能
		if nums[i]+nums[n-2]+nums[n-1] < 0 {
			continue
		}

		//确定统计边界
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[left]+nums[right]+nums[i] == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				//去重第二层
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++

				//去重第三层
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--

			} else if nums[left]+nums[right]+nums[i] > 0 {
				//值太大，需要减小
				right--
			} else {
				//值太小，需要增大
				left++
			}
		}

	}

	return result
}
