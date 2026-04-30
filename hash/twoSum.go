package hash

//1. 两数之和:
//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
//链接:https://leetcode.cn/problems/two-sum/description/?envType=study-plan-v2&envId=top-100-liked

func twoSum(nums []int, target int) []int {
	if nums == nil || len(nums) < 2 {
		return []int{}
	}
	value2Index := make(map[int]int)
	for i, v := range nums {
		j, exist := value2Index[target-v]
		if exist {
			return []int{j, i}
		}

		value2Index[v] = i
	}

	return nil
}
