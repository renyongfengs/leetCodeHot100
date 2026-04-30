package dynamicPlanning

//198. 打家劫舍
//你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，
//影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
//链接：https://leetcode.cn/problems/house-robber/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/house-robber/solutions/2102725/ru-he-xiang-chu-zhuang-tai-ding-yi-he-zh-1wt1/?envType=study-plan-v2&envId=top-100-liked
func rob(nums []int) int {
	memo := make([]int, len(nums))
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(i int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}

		if memo[i] != -1 {
			return memo[i]
		}

		memo[i] = max(dfs(i-1), dfs(i-2)+nums[i])
		return memo[i]
	}

	return dfs(len(nums) - 1)
}

func rob2(nums []int) int {
	memo := make([]int, len(nums)+2)
	for i, num := range nums {
		memo[i+2] = max(memo[i+1], memo[i]+num)
	}
	return memo[len(nums)+1]
}

func rob3(nums []int) int {
	f0, f1 := 0, 0
	for _, num := range nums {
		f1, f0 = max(f0+num, f1), f1
	}
	return f1
}
