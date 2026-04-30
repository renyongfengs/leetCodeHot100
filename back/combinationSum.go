package back

import (
	"slices"
)

//39. 组合总和
//给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，
//并以列表形式返回。你可以按 任意顺序 返回这些组合。
//candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
//对于给定的输入，保证和为 target 的不同组合数少于 150 个。
//链接：https://leetcode.cn/problems/combination-sum/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/combination-sum/solutions/2747858/liang-chong-fang-fa-xuan-huo-bu-xuan-mei-mhf9/?envType=study-plan-v2&envId=top-100-liked
func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	path := []int{}
	var dfs = func(i, target int) {}
	dfs = func(i, target int) {
		//找到
		if target == 0 {
			result = append(result, slices.Clone(path))
			return
		}
		//找不到
		if i == len(candidates) || target < 0 {
			return
		}
		//	该节点的值不选择
		dfs(i+1, target)

		//	该节点的值选择
		path = append(path, candidates[i])
		dfs(i, target-candidates[i])
		path = path[:len(path)-1]
	}

	dfs(0, target)
	return result
}

//剪枝优化
func combinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	result := make([][]int, 0)
	path := []int{}
	var dfs = func(i, target int) {}
	dfs = func(i, target int) {
		//找到
		if target == 0 {
			result = append(result, slices.Clone(path))
			return
		}
		//找不到
		if i == len(candidates) || target < candidates[i] {
			return
		}
		//	该节点的值不选择
		dfs(i+1, target)

		//	该节点的值选择
		path = append(path, candidates[i])
		dfs(i, target-candidates[i])
		path = path[:len(path)-1]
	}

	dfs(0, target)
	return result
}
