package back

import (
	"slices"
)

//78. 子集
//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
//链接：https://leetcode.cn/problems/subsets/description/?envType=study-plan-v2&envId=top-100-liked
func subsets(nums []int) [][]int {
	n := len(nums)
	result := make([][]int, 0)
	path := make([]int, 0, n)
	var dfs func(int)
	dfs = func(i int) {
		result = append(result, slices.Clone(path))
		for j := i; j < n; j++ {
			path = append(path, nums[j])
			dfs(j + 1)
			path = path[:len(path)-1]
		}

	}
	dfs(0)

	return result
}

func subsets2(nums []int) [][]int {
	n := len(nums)
	result := make([][]int, 0)
	path := make([]int, 0, n)

	var dfs func(int)

	dfs = func(i int) {
		if i == n {
			result = append(result, append([]int{}, path...))
			return
		}
		//不选
		dfs(i + 1)

		//选
		path = append(path, nums[i])
		dfs(i + 1)
		//恢复
		path = path[:len(path)-1]
	}
	dfs(0)

	return result
}
