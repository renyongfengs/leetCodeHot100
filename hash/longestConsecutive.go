package hash

import (
	"sort"
)

//128. 最长连续序列
//给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
//链接：https://leetcode.cn/problems/longest-consecutive-sequence/description/?envType=study-plan-v2&envId=top-100-liked

func longestConsecutive(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	vm := make(map[int]bool)
	for _, num := range nums {
		vm[num] = true
	}
	var longestStreak int

	for val := range vm {
		//判断是否是序列的起点
		if !vm[val-1] {

			start, end := val, val+1
			//寻找序列的下一个数字
			for vm[end] {
				end++
			}

			//更新最长序列长度
			if (end - start) > longestStreak {
				longestStreak = end - start
			}

		}
	}

	return longestStreak
}

func longestConsecutive2(nums []int) int {
	if len(nums) == 0 || nums == nil {
		return 0
	}
	//排序
	sort.Ints(nums)
	result := 0
	tempLen := 1
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 1 {
			tempLen++
		} else if nums[i] == nums[i-1] {
			continue
		} else {
			result = max(result, tempLen)
			tempLen = 1
		}

	}

	result = max(result, tempLen)
	return result
}
