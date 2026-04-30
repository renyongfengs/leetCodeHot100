package graph

//994. 腐烂的橘子
//在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：
//值 0 代表空单元格；
//值 1 代表新鲜橘子；
//值 2 代表腐烂的橘子。
//每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。
//返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。
//链接：https://leetcode.cn/problems/rotting-oranges/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/rotting-oranges/solutions/2773461/duo-yuan-bfsfu-ti-dan-pythonjavacgojsrus-yfmh/?envType=study-plan-v2&envId=top-100-liked
func orangesRotting(grid [][]int) int {
	fresh := 0 //新鲜橘子数量
	que := make([]Position, 0)
	m, n := len(grid), len(grid[0])
	for i, row := range grid {
		for j, cell := range row {
			if cell == 1 {
				fresh++
			} else if cell == 2 {
				que = append(que, Position{i, j})
			}
		}
	}
	result := 0
	dfs := func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}

		if grid[i][j] == 1 {
			fresh--
			grid[i][j] = 2
			que = append(que, Position{i, j})
		}
	}

	for len(que) > 0 && fresh > 0 {
		result++

		tmp := que
		que = []Position{}
		for _, q := range tmp {
			dfs(q.x, q.y+1)
			dfs(q.x, q.y-1)
			dfs(q.x+1, q.y)
			dfs(q.x-1, q.y)
		}

	}

	if fresh > 0 {
		return -1
	}

	return result
}

type Position struct {
	x, y int
}
