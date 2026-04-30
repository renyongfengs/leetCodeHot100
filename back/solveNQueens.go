package back

import (
	"strings"
)

//51. N 皇后
//按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
//n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
//每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//链接：https://leetcode.cn/problems/n-queens/?envType=study-plan-v2&envId=top-100-liked
func solveNQueens(n int) (ans [][]string) {
	queens := make([]int, n) // 皇后放在 (r,queens[r])
	col := make([]bool, n)
	diag1 := make([]bool, n*2-1)
	diag2 := make([]bool, n*2-1)
	var dfs func(int)
	dfs = func(r int) {
		if r == n {
			board := make([]string, n)
			for i, c := range queens {
				board[i] = strings.Repeat(".", c) + "Q" + strings.Repeat(".", n-1-c)
			}
			ans = append(ans, board)
			return
		}
		// 在 (r,c) 放皇后
		for c, ok := range col {
			rc := r - c + n - 1
			if !ok && !diag1[r+c] && !diag2[rc] { // 判断能否放皇后
				queens[r] = c                                    // 直接覆盖，无需恢复现场
				col[c], diag1[r+c], diag2[rc] = true, true, true // 皇后占用了 c 列和两条斜线
				dfs(r + 1)
				col[c], diag1[r+c], diag2[rc] = false, false, false // 恢复现场
			}
		}
	}
	dfs(0)
	return
}

func solveNQueens2(n int) (ans [][]string) {
	record := make([][]rune, n)
	for i := range record {
		record[i] = make([]rune, n)
		for j := range record[i] {
			record[i][j] = '.'
		}
	}

	var dfs func(row int)

	dfs = func(row int) {
		if row == n {
			board := make([]string, n)
			for i := range board {
				board[i] = string(record[i])
			}
			ans = append(ans, board)
			return

		}

		for col := 0; col < n; col++ {
			if isInvalid(record, row, col) {
				continue
			}

			record[row][col] = 'Q'
			dfs(row + 1)
			record[row][col] = '.'
		}
	}

	dfs(0)

	return
}

func isInvalid(record [][]rune, row, col int) bool {
	n := len(record)
	// 检查列是否有皇后互相冲突
	for i := 0; i < row; i++ {
		if record[i][col] == 'Q' {
			return true
		}
	}
	// 检查右上方是否有皇后互相冲突
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if record[i][j] == 'Q' {
			return true
		}
	}
	// 检查左上方是否有皇后互相冲突
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if record[i][j] == 'Q' {
			return true
		}
	}
	return false
}
