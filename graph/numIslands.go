package graph

//200. 岛屿数量
//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//此外，你可以假设该网格的四条边均被水包围。
//链接：https://leetcode.cn/problems/number-of-islands/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/number-of-islands/solutions/2965773/ba-fang-wen-guo-de-ge-zi-cha-shang-qi-zi-9gs0/?envType=study-plan-v2&envId=top-100-liked

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int)
	//处理i，j周围的陆地
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
			return
		}

		grid[i][j] = '0'

		dfs(i, j-1) // 往左走
		dfs(i, j+1) // 往右走
		dfs(i-1, j) // 往上走
		dfs(i+1, j) // 往下走
	}
	result := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				result++
			}
		}
	}

	return result
}
