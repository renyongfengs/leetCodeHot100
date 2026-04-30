package dynamicPlanning

//118. 杨辉三角
//给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
//在「杨辉三角」中，每个数是它左上方和右上方的数的和。
//链接：https://leetcode.cn/problems/pascals-triangle/?envType=study-plan-v2&envId=top-100-liked
func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := range result {
		result[i] = make([]int, i+1)

		result[i][0] = 1
		result[i][i] = 1
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j] + result[i-1][j-1]
		}

	}
	return result
}
