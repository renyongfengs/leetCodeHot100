package dynamicPlanning

//300. 最长递增子序列
//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
//链接：https://leetcode.cn/problems/longest-increasing-subsequence/description/?envType=study-plan-v2&envId=top-100-liked
func lengthOfLIS(nums []int) int {
	count := 0
	memo := make([]int, len(nums))
	for _, num := range nums {
		left, right := 0, count
		for left < right {
			mid := left + (right-left)/2
			if memo[mid] >= num {
				right = mid
			} else {
				left = mid + 1
			}
		}

		if left == count {
			count++
		}

		memo[left] = num
	}

	return count
}

func lengthOfLIS2(nums []int) int {
	memo := make([]int, len(nums))
	for i := range memo {
		memo[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				memo[i] = max(memo[i], memo[j]+1)
			}
		}
	}

	result := 0
	for _, v := range memo {
		result = max(result, v)
	}

	return result

}
