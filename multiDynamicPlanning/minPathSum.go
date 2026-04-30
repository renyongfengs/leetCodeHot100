package multiDynamicPlanning

import (
	"math"
)

//64. 最小路径和
//给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//说明：每次只能向下或者向右移动一步。
//链接：https://leetcode.cn/problems/minimum-path-sum/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/minimum-path-sum/solutions/3045828/jiao-ni-yi-bu-bu-si-kao-dpcong-ji-yi-hua-zfb2/?envType=study-plan-v2&envId=top-100-liked
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return math.MaxInt
		}
		if i == 0 && j == 0 {
			return grid[i][j]
		}

		if memo[i][j] == -1 {
			memo[i][j] = min(dfs(i-1, j), dfs(i, j-1)) + grid[i][j]
		}
		return memo[i][j]
	}

	return dfs(m-1, n-1)
}

func minPathSum2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for j := range f[0] {
		f[0][j] = math.MaxInt
	}

	for i, row := range grid {
		f[i+1][0] = math.MaxInt
		for j, col := range row {
			if i == 0 && j == 0 {
				f[1][1] = col
			} else {
				f[i+1][j+1] = min(f[i][j+1], f[i+1][j]) + col
			}
		}
	}

	return f[m][n]
}

func minPathSum3(grid [][]int) int {
	n := len(grid[0])
	f := make([]int, n+1)
	for i := range f {
		f[i] = math.MaxInt
	}
	f[1] = 0
	for _, row := range grid {
		for j, col := range row {
			f[j+1] = min(f[j], f[j+1]) + col
		}
	}

	return f[n]
}
