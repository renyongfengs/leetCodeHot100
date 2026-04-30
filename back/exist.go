package back

import (
	"slices"
)

//79. 单词搜索
//给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
//单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用
//链接：https://leetcode.cn/problems/word-search/description/?envType=study-plan-v2&envId=top-100-liked
func exist(board [][]byte, word string) bool {

	bVM := make(map[byte]int)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			bVM[board[i][j]]++
		}
	}

	wVM := make(map[byte]int)
	w := []byte(word)
	for _, c := range w {
		wVM[byte(c)]++
		if wVM[byte(c)] > bVM[byte(c)] {
			return false
		}
	}

	if wVM[w[0]] > wVM[w[len(w)-1]] {
		slices.Reverse(w)
	}

	positionList := []struct{ x, y int }{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	m, n := len(board), len(board[0])
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k == len(w)-1 {
			return true
		}

		//邻边字符
		board[i][j] = 0
		for _, pos := range positionList {
			x, y := i+pos.x, j+pos.y
			//	判断x,y是否超出边界
			if x >= 0 && x < m && y >= 0 && y < n && dfs(x, y, k+1) {
				return true
			}
		}
		board[i][j] = word[k]
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
