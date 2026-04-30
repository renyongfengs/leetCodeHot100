package matrix

import (
	"slices"
)

//73. 矩阵置零
//给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
//链接：https://leetcode.cn/problems/set-matrix-zeroes/description/?envType=study-plan-v2&envId=top-100-liked
func setZeroes(matrix [][]int) {
	rowZero := make([]bool, len(matrix))
	colZero := make([]bool, len(matrix[0]))

	for i, row := range matrix {
		for j, col := range row {
			if col == 0 {
				rowZero[i] = true
				colZero[j] = true
			}
		}
	}

	for i, row := range rowZero {
		for j, col := range colZero {
			if row || col {
				matrix[i][j] = 0
			}
		}
	}
}

func setZeroes2(matrix [][]int) {
	//把rowZero、colZero记录结果保存在第一行、第一列，for循环从第二行和第二列开始，避免影响第一行、第一列，需要先记录是否有0
	// 记录第一行是否包含 0
	firstRowHasZero := slices.Contains(matrix[0], 0)
	// 记录第一列是否包含 0
	firstColHasZero := false
	for _, row := range matrix {
		if row[0] == 0 {
			firstColHasZero = true
			break
		}
	}
	// 用第一列 matrix[i][0] 保存 rowHasZero[i]
	// 用第一行 matrix[0][j] 保存 colHasZero[j]
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	// 如果第一行一开始就包含 0，那么把第一行全变成 0
	if firstRowHasZero {
		clear(matrix[0])
	}

	// 如果第一列一开始就包含 0，那么把第一列全变成 0
	if firstColHasZero {
		for _, row := range matrix {
			row[0] = 0
		}
	}
}
