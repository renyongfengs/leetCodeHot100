package dynamicPlanning

//416. 分割等和子集
//给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//链接：https://leetcode.cn/problems/partition-equal-subset-sum/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/partition-equal-subset-sum/solutions/553978/bang-ni-ba-0-1bei-bao-xue-ge-tong-tou-by-px33/?envType=study-plan-v2&envId=top-100-liked
// 分割等和子集 动态规划
// 时间复杂度O(n^2) 空间复杂度O(n)
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1)

	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] = max(dp[j], dp[j-num]+num)
		}
	}
	return dp[target] == target
}

func canPartition2(nums []int) bool {
	s := 0
	for _, x := range nums {
		s += x
	}
	if s%2 != 0 {
		return false
	}

	n := len(nums)
	memo := make([][]int8, n)
	for i := range memo {
		memo[i] = make([]int8, s/2+1)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}

	var dfs func(int, int) bool
	dfs = func(i, j int) (res bool) {
		if i < 0 {
			return j == 0
		}
		p := &memo[i][j]
		if *p != -1 { // 之前计算过
			return *p == 1
		}
		if j < nums[i] {
			res = dfs(i-1, j) // 只能不选
		} else {
			res = dfs(i-1, j-nums[i]) || dfs(i-1, j) // 选或不选
		}
		// 记忆化
		if res {
			*p = 1
		} else {
			*p = 0
		}
		return
	}

	return dfs(n-1, s/2)
}
