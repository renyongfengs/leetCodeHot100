package binarySearch

//74. 搜索二维矩阵
//每行中的整数从左到右按非严格递增顺序排列。
//每行的第一个整数大于前一行的最后一个整数。
//给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。
//链接：https://leetcode.cn/problems/search-a-2d-matrix/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/search-a-2d-matrix/solutions/2783931/liang-chong-fang-fa-er-fen-cha-zhao-pai-39d74/?envType=study-plan-v2&envId=top-100-liked
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := -1, m*n
	for left+1 < right {
		mid := left + (right-left)/2
		if matrix[mid/n][mid%n] == target {
			return true
		} else if matrix[mid/n][mid%n] > target {
			right = mid
		} else if matrix[mid/n][mid%n] < target {
			left = mid
		}
	}

	return false
}
