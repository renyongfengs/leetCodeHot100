package matrix

import (
	"slices"
)

//48. 旋转图像
//给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
//你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
//链接：https://leetcode.cn/problems/rotate-image/description/?envType=study-plan-v2&envId=top-100-liked
//讲解：https://leetcode.cn/problems/rotate-image/solutions/3655166/shu-xue-ben-zhi-liang-ci-fan-zhuan-deng-aon4a/?envType=study-plan-v2&envId=top-100-liked
//两次翻转等于一次旋转
//转置：把矩阵按照主对角线翻转，位于 (i,j) 的元素去到 (j,i)。
//行翻转：把每一行翻转，位于 (j,i) 的元素去到 (j,n−1−i)。
func rotate(matrix [][]int) {
	//矩阵转置
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	//矩阵行翻转
	for _, rowElems := range matrix {
		slices.Reverse(rowElems)
	}
}

//遍历顺序调整为遍历对角线上方元素，这样每行遍历完后，这一行的元素后面不会再访问到，可以直接做行翻转。
func rotate2(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
		slices.Reverse(matrix[i])
	}
}
