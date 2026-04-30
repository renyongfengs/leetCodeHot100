package multiDynamicPlanning

//62. 不同路径
//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//问总共有多少条不同的路径？
//链接：https://leetcode.cn/problems/unique-paths/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/unique-paths/solutions/3062432/liang-chong-fang-fa-dong-tai-gui-hua-zu-o5k32/?envType=study-plan-v2&envId=top-100-liked
//递归
func uniquePaths(m int, n int) int {
	memo := make([][]int, m)
	for i := range memo {
		tmp := make([]int, n)
		for j := range tmp {
			tmp[j] = -1
		}
		memo[i] = tmp
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}

		if i == 0 && j == 0 {
			return 1
		}

		if memo[i][j] == -1 {
			memo[i][j] = dfs(i, j-1) + dfs(i-1, j)
		}

		return memo[i][j]
	}
	return dfs(m-1, n-1)
}

//动态规划
func uniquePaths2(m int, n int) int {
	memo := make([][]int, m+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}
	memo[0][1] = 1
	for i := range m {
		for j := range n {
			memo[i+1][j+1] = memo[i][j+1] + memo[i+1][j]
		}
	}
	return memo[m][n]
}

func uniquePaths3(m int, n int) int {
	f := make([]int, n+1)
	f[1] = 1
	for range m {
		for j := range n {
			f[j+1] += f[j]
		}
	}
	return f[n]
}
